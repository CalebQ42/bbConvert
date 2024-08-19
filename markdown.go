package bbConvert

import (
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

type MarkdownConverter struct {
	largeCodeConv  *regexp2.Regexp
	inlineCodeConv *regexp2.Regexp
	surroundConv   *regexp2.Regexp
	linkImgConv    *regexp2.Regexp
	blockQuoteConv *regexp2.Regexp
	bulletConv     *regexp2.Regexp
	numListConv    *regexp2.Regexp
	headingConv    *regexp2.Regexp
}

func NewMarkdownConverter() MarkdownConverter {
	largeCodeConv := regexp2.MustCompile("```([\\s\\S])```", regexp2.Multiline)
	inlineCodeConv := regexp2.MustCompile("`(.*?)`", regexp2.None)
	bulletConv := regexp2.MustCompile(`(?<!.)([ \t]*)[\*-] (.*)\n`, regexp2.None)
	numListConv := regexp2.MustCompile(`(?<!.)([ \t]*)[0-9]+[.)] (.*)\n`, regexp2.None)
	surroundConv := regexp2.MustCompile(`(\*\*\*|\*\*|\*|___|__|_|~~)(.*?)(\1)`, regexp2.None)
	linkImgConv := regexp2.MustCompile(`[!]?\[(.*?)\]\((.*?)\)`, regexp2.None)
	blockQuoteConv := regexp2.MustCompile(`(?<!.)>(.*)`, regexp2.None)
	headingConv := regexp2.MustCompile(`(?<!.)(#+?) (.*)`, regexp2.None)
	return MarkdownConverter{
		largeCodeConv:  largeCodeConv,
		inlineCodeConv: inlineCodeConv,
		bulletConv:     bulletConv,
		numListConv:    numListConv,
		surroundConv:   surroundConv,
		linkImgConv:    linkImgConv,
		blockQuoteConv: blockQuoteConv,
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
	// Bullet points (unordered list)
	for {
		var allMatches []*regexp2.Match
		var prevMatch *regexp2.Match
		// Find all lines that are in a single list
		for {
			var newMatch *regexp2.Match
			newMatch, err = m.bulletConv.FindNextMatch(prevMatch)
			if newMatch == nil || err != nil {
				break
			}
			if prevMatch != nil {
				if newMatch.Index != prevMatch.Index+prevMatch.Length {
					break
				}
			}
			allMatches = append(allMatches, newMatch)
			prevMatch = newMatch
		}
		if len(allMatches) == 0 {
			break
		}
		converted := "<ul>"
		curHeight := calculateListLevel(allMatches[0].GroupByNumber(1).String())
		curLvl := 1
		for _, m := range allMatches {
			itemHeight := calculateListLevel(m.GroupByNumber(1).String())
			if itemHeight > curHeight {
				curLvl++
				curHeight = itemHeight
				converted += "<ul>"
			} else if itemHeight < curHeight {
				if curLvl > 1 {
					curLvl--
					curHeight = itemHeight
					converted += "</ul>"
				}
			}
			converted += "<li>" + m.GroupByNumber(2).String() + "</li>"
		}
		converted += strings.Repeat("</ul>", curLvl)
	}
	// Numbered list
	for {
		var allMatches []*regexp2.Match
		var prevMatch *regexp2.Match
		// Find all lines that are in a single list
		for {
			var newMatch *regexp2.Match
			newMatch, err = m.numListConv.FindNextMatch(prevMatch)
			if newMatch == nil || err != nil {
				break
			}
			if prevMatch != nil {
				if newMatch.Index != prevMatch.Index+prevMatch.Length {
					break
				}
			}
			allMatches = append(allMatches, newMatch)
			prevMatch = newMatch
		}
		if len(allMatches) == 0 {
			break
		}
		converted := "<ol>"
		curHeight := calculateListLevel(allMatches[0].GroupByNumber(1).String())
		curLvl := 1
		for _, m := range allMatches {
			itemHeight := calculateListLevel(m.GroupByNumber(1).String())
			if itemHeight > curHeight {
				curLvl++
				curHeight = itemHeight
				converted += "<ol>"
			} else if itemHeight < curHeight {
				if curLvl > 1 {
					curLvl--
					curHeight = itemHeight
					converted += "</ol>"
				}
			}
			converted += "<li>" + m.GroupByNumber(2).String() + "</li>"
		}
		converted += strings.Repeat("</ol>", curLvl)
		in = in[:allMatches[0].Index] + converted + in[prevMatch.Index+prevMatch.Length:]
	}
	// Surround (eg. *hi*)
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
		case "**":
			fallthrough
		case "__":
			converted = "<b>" + match.GroupByNumber(2).String() + "</b>"
		case "~~":
			converted = "<s>" + match.GroupByNumber(2).String() + "</s>"
		case "***":
			fallthrough
		case "___":
			converted = "<b><i>" + match.GroupByNumber(2).String() + "</i></b>"
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
				strings.ReplaceAll(match.GroupByNumber(1).String(), "'", "\\'") + ">"
		} else {
			converted = "<a href='" +
				strings.ReplaceAll(match.GroupByNumber(2).String(), "'", "\\'") + ">" +
				match.GroupByNumber(1).String() + "</a>"
		}
		in = in[:match.Index] + converted + in[match.Index+match.Length:]
	}
	// Block Quotes
	for {
		match, err = m.blockQuoteConv.FindStringMatch(in)
		if err != nil || match == nil {
			break
		}
		in = in[:match.Index] + "<blockquote>" + match.GroupByNumber(1).String() + "</blockquote>" + in[match.Index+match.Length:]
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
	return len(indent) % 2
}
