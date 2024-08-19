package bbConvert

import (
	"slices"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

type MarkdownConverter struct {
	largeCodeConv  *regexp2.Regexp
	inlineCodeConv *regexp2.Regexp
	blockQuoteConv *regexp2.Regexp
	bAndIConv      *regexp2.Regexp
	bConv          *regexp2.Regexp
	surroundConv   *regexp2.Regexp
	linkImgConv    *regexp2.Regexp
	listConv       *regexp2.Regexp
	headingConv    *regexp2.Regexp
	//TODO: Tables???
}

func NewMarkdownConverter() MarkdownConverter {
	largeCodeConv := regexp2.MustCompile("```([\\s\\S]*)```", regexp2.Multiline)
	inlineCodeConv := regexp2.MustCompile("`(.*)`", regexp2.None)
	listConv := regexp2.MustCompile(`(?<!.+)([ \t]*)([\*-]|(?:[0-9]+[.)])) (.*)\n`, regexp2.None)
	blockQuoteConv := regexp2.MustCompile(`(?<!.+)(>+) ?(.*)\n`, regexp2.None)
	bAndIConv := regexp2.MustCompile(`(\*\*\*|___)(.*?)(\1)`, regexp2.None)
	bConv := regexp2.MustCompile(`(\*\*|__)(.*?)(\1)`, regexp2.None)
	surroundConv := regexp2.MustCompile(`(\*|_|~~)(.*?)(\1)`, regexp2.None)
	linkImgConv := regexp2.MustCompile(`[!]?\[(.*?)\]\((.*?)\)`, regexp2.None)
	headingConv := regexp2.MustCompile(`(?<!.+)(#+) (.*)`, regexp2.None)
	return MarkdownConverter{
		largeCodeConv:  largeCodeConv,
		inlineCodeConv: inlineCodeConv,
		listConv:       listConv,
		blockQuoteConv: blockQuoteConv,
		bAndIConv:      bAndIConv,
		bConv:          bConv,
		surroundConv:   surroundConv,
		linkImgConv:    linkImgConv,
		headingConv:    headingConv,
	}
}

func (m MarkdownConverter) HTMLConvert(in string) string {
	var codeBlocks []string
	var match *regexp2.Match
	var err error
	// Code blocks
	for {
		match, err = m.largeCodeConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		codeBlocks = append(codeBlocks, match.GroupByNumber(1).String())
		in = in[:match.Index] + codePlaceholder + in[match.Index+match.Length:]
	}
	// Inline code
	for {
		match, err = m.inlineCodeConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		codeBlocks = append(codeBlocks, match.GroupByNumber(1).String())
		in = in[:match.Index] + codePlaceholder + in[match.Index+match.Length:]
	}
	// lists (ordered and unordered)
	for {
		var allMatches []*regexp2.Match
		var prevMatch *regexp2.Match
		prevMatch, err = m.listConv.FindStringMatch(in)
		if err != nil || prevMatch == nil {
			break
		}
		allMatches = append(allMatches, prevMatch)
		// Find all lines that are in a single list
		for {
			var newMatch *regexp2.Match
			newMatch, err = m.listConv.FindNextMatch(prevMatch)
			if newMatch == nil || err != nil {
				break
			}
			if newMatch.Index != prevMatch.Index+prevMatch.Length {
				break
			}
			allMatches = append(allMatches, newMatch)
			prevMatch = newMatch
		}
		if len(allMatches) == 0 {
			break
		}
		var isOrdered []bool
		var converted string
		if allMatches[0].GroupByNumber(2).String() == "*" || allMatches[0].GroupByNumber(2).String() == "-" {
			converted = "<ul>"
			isOrdered = append(isOrdered, false)
		} else {
			converted = "<ol>"
			isOrdered = append(isOrdered, true)
		}
		curHeight := calculateListLevel(allMatches[0].GroupByNumber(1).String())
		curLvl := 1
		for _, m := range allMatches {
			itemHeight := calculateListLevel(m.GroupByNumber(1).String())
			if itemHeight > curHeight {
				curLvl++
				curHeight = itemHeight
				if m.GroupByNumber(2).String() == "*" || m.GroupByNumber(2).String() == "-" {
					converted += "<ul>"
					isOrdered = append(isOrdered, false)
				} else {
					converted += "<ol>"
					isOrdered = append(isOrdered, true)
				}
			} else if itemHeight < curHeight {
				if curLvl > 1 {
					curLvl--
					curHeight = itemHeight
					if isOrdered[len(isOrdered)-1] {
						converted += "</ol>"
					} else {
						converted += "</ul>"
					}
					isOrdered = isOrdered[:len(isOrdered)-1]
					if m.GroupByNumber(2).String() == "*" || m.GroupByNumber(2).String() == "-" {
						if isOrdered[len(isOrdered)-1] {
							curLvl++
							converted += "<ul>"
							isOrdered = append(isOrdered, false)
						}
					} else {
						if !isOrdered[len(isOrdered)-1] {
							curLvl++
							converted += "<ol>"
							isOrdered = append(isOrdered, true)
						}
					}
				}
			}
			converted += "<li>" + m.GroupByNumber(3).String() + "</li>"
		}
		for _, b := range slices.Backward(isOrdered) {
			if b {
				converted += "</ol>"
			} else {
				converted += "</ul>"
			}
		}
		in = in[:allMatches[0].Index] + converted + in[prevMatch.Index+prevMatch.Length:]
	}
	// Block Quotes
	for {
		var allMatches []*regexp2.Match
		var prevMatch *regexp2.Match
		prevMatch, err = m.blockQuoteConv.FindStringMatch(in)
		if err != nil || prevMatch == nil {
			break
		}
		allMatches = append(allMatches, prevMatch)
		// Find all lines that are in a single blockquote
		for {
			var newMatch *regexp2.Match
			newMatch, err = m.blockQuoteConv.FindNextMatch(prevMatch)
			if newMatch == nil || err != nil {
				break
			}
			if newMatch.Index != prevMatch.Index+prevMatch.Length {
				break
			}
			allMatches = append(allMatches, newMatch)
			prevMatch = newMatch
		}
		if len(allMatches) == 0 {
			break
		}
		curHeight := len(allMatches[0].GroupByNumber(1).String())
		curLvl := 1
		converted := "<blockquote><p>"
		for _, m := range allMatches {
			if m.GroupByNumber(2).String() == "" {
				converted += "</p><p>"
				continue
			}
			itemHeight := len(m.GroupByNumber(1).String())
			if itemHeight > curHeight {
				curHeight = itemHeight
				curLvl++
				converted += "</p><blockquote><p>"
			} else if itemHeight < curHeight && curLvl > 1 {
				curHeight = itemHeight
				curLvl--
				converted += "</p></blockquote><p>"
			}
			converted += m.GroupByNumber(2).String()
		}
		converted += strings.Repeat("</p></blockquote>", curLvl)
		in = in[:allMatches[0].Index] + converted + in[prevMatch.Index+prevMatch.Length:]
	}
	// Bold and Italics (*** and ___)
	for {
		match, err = m.bAndIConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		in = in[:match.Index] + "<b><i>" + match.GroupByNumber(2).String() + "</i></b>" + in[match.Index+match.Length:]
	}
	// Bold (** and __)
	for {
		match, err = m.bConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		in = in[:match.Index] + "<b>" + match.GroupByNumber(2).String() + "</b>" + in[match.Index+match.Length:]
	}
	// Italics and strikethough
	for {
		match, err = m.surroundConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		var converted string
		switch match.GroupByNumber(1).String() {
		case "*":
			fallthrough
		case "_":
			converted = "<i>" + match.GroupByNumber(2).String() + "</i>"
		case "~~":
			converted = "<s>" + match.GroupByNumber(2).String() + "</s>"
		default:
			converted = match.GroupByNumber(2).String()
		}
		in = in[:match.Index] + converted + in[match.Index+match.Length:]
	}
	// Links and images
	for {
		match, err = m.linkImgConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		var converted string
		if match.String()[0] == '!' {
			converted = "<img src='" +
				strings.ReplaceAll(match.GroupByNumber(2).String(), "'", "\\'") + "' alt='" +
				strings.ReplaceAll(match.GroupByNumber(1).String(), "'", "\\'") + "'>"
		} else {
			converted = "<a href='" +
				strings.ReplaceAll(match.GroupByNumber(2).String(), "'", "\\'") + "'>" +
				match.GroupByNumber(1).String() + "</a>"
		}
		in = in[:match.Index] + converted + in[match.Index+match.Length:]
	}
	// Headings
	for {
		match, err = m.headingConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		level := len(match.GroupByNumber(1).String())
		if level > 6 {
			level = 6
		}
		in = in[:match.Index] +
			"<h" + strconv.Itoa(level) + ">" +
			match.GroupByNumber(2).String() +
			"</h" + strconv.Itoa(level) + ">" +
			in[match.Index+match.Length:]
	}
	in = "<p>" + strings.ReplaceAll(in, "\n\n", "</p>\n<p>") + "</p>"
	// Replace the code placeholders
	for i := range codeBlocks {
		if strings.Contains(codeBlocks[i], "\n") {
			in = strings.Replace(in, codePlaceholder, "<pre><code>"+codeBlocks[i]+"</code></pre>", 1)
		} else {
			in = strings.Replace(in, codePlaceholder, "<code>"+codeBlocks[i]+"</code>", 1)
		}
	}
	return in
}

func calculateListLevel(indent string) int {
	indent = strings.ReplaceAll(indent, "\t", "  ")
	return len(indent) / 2
}
