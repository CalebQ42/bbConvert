package bbConvert

import (
	"github.com/dlclark/regexp2"
	"github.com/dlclark/regexp2/helpers"
	"github.com/dlclark/regexp2/syntax"
	"unicode"
)

/*
Capture-M(index = 0, unindex = -1)
 Concatenate-M
  Multi-M(String = "```")
  Capture-M(index = 1, unindex = -1)
   Set-M(Set = [\x00-\u10ffff])
  Multi-M(String = "```")
*/
// From markdown.go:22:39
// Pattern: "```([\\s\\S])```"
// Options: regexp2.Multiline
type largeCodeConv_Engine struct{}

func (largeCodeConv_Engine) Caps() map[int]int        { return nil }
func (largeCodeConv_Engine) CapNames() map[string]int { return nil }
func (largeCodeConv_Engine) CapsList() []string       { return nil }
func (largeCodeConv_Engine) CapSize() int             { return 2 }

func (largeCodeConv_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 7 characters
	if pos <= len(r.Runtext)-7 {
		// The pattern has the literal "```" at the beginning of the pattern. Find the next occurrence.
		// If it can't be found, there's no match
		if i := helpers.IndexOf(r.Runtext[pos:], []rune("```")); i >= 0 {
			r.Runtextpos = pos + i
			return true
		}
	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (largeCodeConv_Engine) Execute(r *regexp2.Runner) error {
	capture_starting_pos := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Concatenate-M
	// Node: Multi-M(String = "```")
	// Match the string "```".
	if !helpers.StartsWith(slice, []rune("```")) {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	// Node: Capture-M(index = 1, unindex = -1)
	// "1" capture group
	pos += 3
	slice = r.Runtext[pos:]
	capture_starting_pos = pos

	// Node: Set-M(Set = [\x00-\u10ffff])
	// Match [\x00-\u10ffff].
	if len(slice) == 0 || false {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	pos++
	slice = r.Runtext[pos:]
	r.Capture(1, capture_starting_pos, pos)

	// Node: Multi-M(String = "```")
	// Match the string "```".
	if !helpers.StartsWith(slice, []rune("```")) {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	// The input matched.
	pos += 3
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture-M(index = 0, unindex = -1)
 Concatenate-M
  Multi-M(String = "[code]")
  Capture-M(index = 1, unindex = -1)
   Setlazy-M(Set = [\x00-\u10ffff])(Min = 0, Max = inf)
  Multi-M(String = "[/code]")
*/
// From bb.go:37:30
// Pattern: "\\[code\\]([\\s\\S]*?)\\[\\/code\\]"
// Options: regexp2.Multiline
type code_Engine struct{}

func (code_Engine) Caps() map[int]int        { return nil }
func (code_Engine) CapNames() map[string]int { return nil }
func (code_Engine) CapsList() []string       { return nil }
func (code_Engine) CapSize() int             { return 2 }

func (code_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 13 characters
	if pos <= len(r.Runtext)-13 {
		// The pattern has the literal "[code]" at the beginning of the pattern. Find the next occurrence.
		// If it can't be found, there's no match
		if i := helpers.IndexOf(r.Runtext[pos:], []rune("[code]")); i >= 0 {
			r.Runtextpos = pos + i
			return true
		}
	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (code_Engine) Execute(r *regexp2.Runner) error {
	capture_starting_pos := 0
	lazyloop_capturepos := 0
	lazyloop_pos := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Concatenate-M
	// Node: Multi-M(String = "[code]")
	// Match the string "[code]".
	if !helpers.StartsWith(slice, []rune("[code]")) {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	// Node: Capture-M(index = 1, unindex = -1)
	// "1" capture group
	pos += 6
	slice = r.Runtext[pos:]
	capture_starting_pos = pos

	// Node: Setlazy-M(Set = [\x00-\u10ffff])(Min = 0, Max = inf)
	// Match [\x00-\u10ffff] lazily any number of times.
	lazyloop_pos = pos
	goto LazyLoopEnd

LazyLoopBacktrack:
	r.UncaptureUntil(lazyloop_capturepos)
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	pos = lazyloop_pos
	slice = r.Runtext[pos:]
	if len(slice) == 0 || false {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}
	pos++
	slice = r.Runtext[pos:]
	lazyloop_pos = helpers.IndexOf(slice, []rune("[/code]"))
	if lazyloop_pos < 0 {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}
	pos += lazyloop_pos
	slice = r.Runtext[pos:]
	lazyloop_pos = pos

LazyLoopEnd:
	lazyloop_capturepos = r.Crawlpos()

	r.Capture(1, capture_starting_pos, pos)

	goto CaptureSkipBacktrack

CaptureBacktrack:
	goto LazyLoopBacktrack

CaptureSkipBacktrack:
	;

	// Node: Multi-M(String = "[/code]")
	// Match the string "[/code]".
	if !helpers.StartsWith(slice, []rune("[/code]")) {
		goto CaptureBacktrack
	}

	// The input matched.
	pos += 7
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture-M(index = 0, unindex = -1)
 Concatenate-M
  One-M(Ch = \[)
  Capture-M(index = 1, unindex = -1)
   Alternate-M
    Concatenate-M
     One-M(Ch = b)
     Alternate-M
      Empty-M
      Multi-M(String = "old")
    Concatenate-M
     One-M(Ch = i)
     Alternate-M
      Empty-M
      Multi-M(String = "talics")
    Concatenate-M
     One-M(Ch = u)
     Alternate-M
      Empty-M
      Multi-M(String = "nderline")
    Concatenate-M
     One-M(Ch = s)
     Alternate-M
      Empty-M
      Multi-M(String = "trike")
    Multi-M(String = "font")
    Multi-M(String = "size")
    Multi-M(String = "color")
    Multi-M(String = "smallcaps")
    Multi-M(String = "url")
    Multi-M(String = "link")
    Multi-M(String = "youtube")
    Concatenate-M
     Multi-M(String = "im")
     Alternate-M
      One-M(Ch = g)
      Multi-M(String = "age")
    Concatenate-M
     One-M(Ch = t)
     Alternate-M
      Multi-M(String = "itle")
      Set-M(Set = [1-6])
    Multi-M(String = "align")
    Multi-M(String = "float")
    Multi-M(String = "ul")
    Multi-M(String = "bullet")
    Multi-M(String = "ol")
    Multi-M(String = "number")
  Capture-M(index = 2, unindex = -1)
   Notonelazy-M(Ch = \n)(Min = 0, Max = inf)
  One-M(Ch = \])
  Capture-M(index = 3, unindex = -1)
   Setlazy-M(Set = [\x00-\u10ffff])(Min = 0, Max = inf)
  Multi-M(String = "[/")
  Ref-M(index = 1)
  One-M(Ch = \])
*/
// From bb.go:38:30
// Pattern: "\\[(b|bold|i|italics|u|underline|s|strike|font|size|color|smallcaps|url|link|youtube|img|image|title|t[1-6]|align|float|ul|bullet|ol|number)(.*?)\\]([\\s\\S]*?)\\[\\/\\1\\]"
// Options: regexp2.Multiline
type main_Engine struct{}

func (main_Engine) Caps() map[int]int        { return nil }
func (main_Engine) CapNames() map[string]int { return nil }
func (main_Engine) CapsList() []string       { return nil }
func (main_Engine) CapSize() int             { return 4 }

func (main_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 6 characters
	if pos <= len(r.Runtext)-6 {
		// The pattern begins with [\[]
		// Find the next occurrence. If it can't be found, there's no match.
		i := helpers.IndexOfAny1(r.Runtext[pos:], '[')
		if i >= 0 {
			r.Runtextpos = pos + i
			return true
		}

	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (main_Engine) Execute(r *regexp2.Runner) error {
	capture_starting_pos := 0
	alternation_starting_pos := 0
	alternation_starting_capturepos := 0
	alternation_branch := 0
	alternation_starting_pos1 := 0
	alternation_starting_capturepos1 := 0
	alternation_branch1 := 0
	alternation_starting_pos2 := 0
	alternation_starting_capturepos2 := 0
	alternation_branch2 := 0
	alternation_starting_pos3 := 0
	alternation_starting_capturepos3 := 0
	alternation_branch3 := 0
	alternation_starting_pos4 := 0
	alternation_starting_capturepos4 := 0
	alternation_branch4 := 0
	capture_starting_pos1 := 0
	lazyloop_capturepos := 0
	lazyloop_pos := 0
	capture_starting_pos2 := 0
	lazyloop_capturepos1 := 0
	lazyloop_pos1 := 0
	matchLength := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Concatenate-M
	// Node: One-M(Ch = \[)
	// Match '['.
	if len(slice) == 0 || slice[0] != '[' {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	// Node: Capture-M(index = 1, unindex = -1)
	// "1" capture group
	pos++
	slice = r.Runtext[pos:]
	capture_starting_pos = pos

	// Node: Alternate-M
	// Match with 19 alternative expressions.
	alternation_starting_pos = pos
	alternation_starting_capturepos = r.Crawlpos()

	// Branch 0
	// Node: Concatenate-M
	// Node: One-M(Ch = b)
	// Match 'b'.
	if len(slice) == 0 || slice[0] != 'b' {
		goto AlternationBranch
	}

	// Node: Alternate-M
	// Match with 2 alternative expressions.
	alternation_starting_pos1 = pos
	alternation_starting_capturepos1 = r.Crawlpos()

	// Branch 0
	// Node: Empty-M

	alternation_branch1 = 0
	pos++
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch1:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 1
	// Node: Multi-M(String = "old")
	// Match the string "old".
	if !helpers.StartsWith(slice[1:], []rune("old")) {
		goto AlternationBranch
	}

	alternation_branch1 = 1
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBacktrack1:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	switch alternation_branch1 {
	case 0:
		goto AlternationBranch1
	case 1:
		goto AlternationBranch
	}

AlternationMatch1:
	;

	alternation_branch = 0
	goto AlternationMatch

AlternationBranch:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 1
	// Node: Concatenate-M
	// Node: One-M(Ch = i)
	// Match 'i'.
	if len(slice) == 0 || slice[0] != 'i' {
		goto AlternationBranch2
	}

	// Node: Alternate-M
	// Match with 2 alternative expressions.
	alternation_starting_pos2 = pos
	alternation_starting_capturepos2 = r.Crawlpos()

	// Branch 0
	// Node: Empty-M

	alternation_branch2 = 0
	pos++
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch3:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 1
	// Node: Multi-M(String = "talics")
	// Match the string "talics".
	if !helpers.StartsWith(slice[1:], []rune("talics")) {
		goto AlternationBranch2
	}

	alternation_branch2 = 1
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBacktrack2:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	switch alternation_branch2 {
	case 0:
		goto AlternationBranch3
	case 1:
		goto AlternationBranch2
	}

AlternationMatch2:
	;

	alternation_branch = 1
	goto AlternationMatch

AlternationBranch2:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 2
	// Node: Concatenate-M
	// Node: One-M(Ch = u)
	// Match 'u'.
	if len(slice) == 0 || slice[0] != 'u' {
		goto AlternationBranch4
	}

	// Node: Alternate-M
	// Match with 2 alternative expressions.
	alternation_starting_pos3 = pos
	alternation_starting_capturepos3 = r.Crawlpos()

	// Branch 0
	// Node: Empty-M

	alternation_branch3 = 0
	pos++
	slice = r.Runtext[pos:]
	goto AlternationMatch3

AlternationBranch5:
	pos = alternation_starting_pos3
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos3)

	// Branch 1
	// Node: Multi-M(String = "nderline")
	// Match the string "nderline".
	if !helpers.StartsWith(slice[1:], []rune("nderline")) {
		goto AlternationBranch4
	}

	alternation_branch3 = 1
	pos += 9
	slice = r.Runtext[pos:]
	goto AlternationMatch3

AlternationBacktrack3:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	switch alternation_branch3 {
	case 0:
		goto AlternationBranch5
	case 1:
		goto AlternationBranch4
	}

AlternationMatch3:
	;

	alternation_branch = 2
	goto AlternationMatch

AlternationBranch4:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 3
	// Node: Concatenate-M
	// Node: One-M(Ch = s)
	// Match 's'.
	if len(slice) == 0 || slice[0] != 's' {
		goto AlternationBranch6
	}

	// Node: Alternate-M
	// Match with 2 alternative expressions.
	alternation_starting_pos4 = pos
	alternation_starting_capturepos4 = r.Crawlpos()

	// Branch 0
	// Node: Empty-M

	alternation_branch4 = 0
	pos++
	slice = r.Runtext[pos:]
	goto AlternationMatch4

AlternationBranch7:
	pos = alternation_starting_pos4
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos4)

	// Branch 1
	// Node: Multi-M(String = "trike")
	// Match the string "trike".
	if !helpers.StartsWith(slice[1:], []rune("trike")) {
		goto AlternationBranch6
	}

	alternation_branch4 = 1
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch4

AlternationBacktrack4:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	switch alternation_branch4 {
	case 0:
		goto AlternationBranch7
	case 1:
		goto AlternationBranch6
	}

AlternationMatch4:
	;

	alternation_branch = 3
	goto AlternationMatch

AlternationBranch6:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 4
	// Node: Multi-M(String = "font")
	// Match the string "font".
	if !helpers.StartsWith(slice, []rune("font")) {
		goto AlternationBranch8
	}

	alternation_branch = 4
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch8:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 5
	// Node: Multi-M(String = "size")
	// Match the string "size".
	if !helpers.StartsWith(slice, []rune("size")) {
		goto AlternationBranch9
	}

	alternation_branch = 5
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch9:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 6
	// Node: Multi-M(String = "color")
	// Match the string "color".
	if !helpers.StartsWith(slice, []rune("color")) {
		goto AlternationBranch10
	}

	alternation_branch = 6
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch10:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 7
	// Node: Multi-M(String = "smallcaps")
	// Match the string "smallcaps".
	if !helpers.StartsWith(slice, []rune("smallcaps")) {
		goto AlternationBranch11
	}

	alternation_branch = 7
	pos += 9
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch11:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 8
	// Node: Multi-M(String = "url")
	// Match the string "url".
	if !helpers.StartsWith(slice, []rune("url")) {
		goto AlternationBranch12
	}

	alternation_branch = 8
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch12:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 9
	// Node: Multi-M(String = "link")
	// Match the string "link".
	if !helpers.StartsWith(slice, []rune("link")) {
		goto AlternationBranch13
	}

	alternation_branch = 9
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch13:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 10
	// Node: Multi-M(String = "youtube")
	// Match the string "youtube".
	if !helpers.StartsWith(slice, []rune("youtube")) {
		goto AlternationBranch14
	}

	alternation_branch = 10
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch14:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 11
	// Node: Concatenate-M
	// Node: Multi-M(String = "im")
	// Match the string "im".
	if !helpers.StartsWith(slice, []rune("im")) {
		goto AlternationBranch15
	}

	// Node: Alternate-M
	// Match with 2 alternative expressions.
	if len(slice) < 3 {
		goto AlternationBranch15
	}

	switch slice[2] {
	case 'g':
		pos += 3
		slice = r.Runtext[pos:]

	case 'a':
		// Node: Multi-M(String = "ge")
		// Match the string "ge".
		if !helpers.StartsWith(slice[3:], []rune("ge")) {
			goto AlternationBranch15
		}

		pos += 5
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch15
	}

	alternation_branch = 11
	goto AlternationMatch

AlternationBranch15:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 12
	// Node: Concatenate-M
	// Node: One-M(Ch = t)
	// Match 't'.
	if len(slice) == 0 || slice[0] != 't' {
		goto AlternationBranch16
	}

	// Node: Alternate-M
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch16
	}

	switch slice[1] {
	case 'i':
		// Node: Multi-M(String = "tle")
		// Match the string "tle".
		if !helpers.StartsWith(slice[2:], []rune("tle")) {
			goto AlternationBranch16
		}

		pos += 5
		slice = r.Runtext[pos:]

	case '1', '2', '3', '4', '5', '6':
		pos += 2
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch16
	}

	alternation_branch = 12
	goto AlternationMatch

AlternationBranch16:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 13
	// Node: Multi-M(String = "align")
	// Match the string "align".
	if !helpers.StartsWith(slice, []rune("align")) {
		goto AlternationBranch17
	}

	alternation_branch = 13
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch17:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 14
	// Node: Multi-M(String = "float")
	// Match the string "float".
	if !helpers.StartsWith(slice, []rune("float")) {
		goto AlternationBranch18
	}

	alternation_branch = 14
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch18:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 15
	// Node: Multi-M(String = "ul")
	// Match the string "ul".
	if !helpers.StartsWith(slice, []rune("ul")) {
		goto AlternationBranch19
	}

	alternation_branch = 15
	pos += 2
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch19:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 16
	// Node: Multi-M(String = "bullet")
	// Match the string "bullet".
	if !helpers.StartsWith(slice, []rune("bullet")) {
		goto AlternationBranch20
	}

	alternation_branch = 16
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch20:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 17
	// Node: Multi-M(String = "ol")
	// Match the string "ol".
	if !helpers.StartsWith(slice, []rune("ol")) {
		goto AlternationBranch21
	}

	alternation_branch = 17
	pos += 2
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBranch21:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 18
	// Node: Multi-M(String = "number")
	// Match the string "number".
	if !helpers.StartsWith(slice, []rune("number")) {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	alternation_branch = 18
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBacktrack:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	switch alternation_branch {
	case 0:
		goto AlternationBacktrack1
	case 1:
		goto AlternationBacktrack2
	case 2:
		goto AlternationBacktrack3
	case 3:
		goto AlternationBacktrack4
	case 4:
		goto AlternationBranch8
	case 5:
		goto AlternationBranch9
	case 6:
		goto AlternationBranch10
	case 7:
		goto AlternationBranch11
	case 8:
		goto AlternationBranch12
	case 9:
		goto AlternationBranch13
	case 10:
		goto AlternationBranch14
	case 11:
		goto AlternationBranch15
	case 12:
		goto AlternationBranch16
	case 13:
		goto AlternationBranch17
	case 14:
		goto AlternationBranch18
	case 15:
		goto AlternationBranch19
	case 16:
		goto AlternationBranch20
	case 17:
		goto AlternationBranch21
	case 18:
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

AlternationMatch:
	;

	r.Capture(1, capture_starting_pos, pos)

	goto CaptureSkipBacktrack

CaptureBacktrack:
	goto AlternationBacktrack

CaptureSkipBacktrack:
	;

	// Node: Capture-M(index = 2, unindex = -1)
	// "2" capture group
	capture_starting_pos1 = pos

	// Node: Notonelazy-M(Ch = \n)(Min = 0, Max = inf)
	// Match a character other than '\n' lazily any number of times.
	lazyloop_pos = pos
	goto LazyLoopEnd

LazyLoopBacktrack:
	r.UncaptureUntil(lazyloop_capturepos)
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	pos = lazyloop_pos
	slice = r.Runtext[pos:]
	if len(slice) == 0 || slice[0] == '\n' {
		goto CaptureBacktrack
	}
	pos++
	slice = r.Runtext[pos:]
	lazyloop_pos = helpers.IndexOfAny2(slice, '\n', ']')
	if lazyloop_pos >= len(slice) || slice[lazyloop_pos] == '\n' {
		goto CaptureBacktrack
	}
	pos += lazyloop_pos
	slice = r.Runtext[pos:]
	lazyloop_pos = pos

LazyLoopEnd:
	lazyloop_capturepos = r.Crawlpos()

	r.Capture(2, capture_starting_pos1, pos)

	goto CaptureSkipBacktrack1

CaptureBacktrack1:
	goto LazyLoopBacktrack

CaptureSkipBacktrack1:
	;

	// Node: One-M(Ch = \])
	// Match ']'.
	if len(slice) == 0 || slice[0] != ']' {
		goto CaptureBacktrack1
	}

	// Node: Capture-M(index = 3, unindex = -1)
	// "3" capture group
	pos++
	slice = r.Runtext[pos:]
	capture_starting_pos2 = pos

	// Node: Setlazy-M(Set = [\x00-\u10ffff])(Min = 0, Max = inf)
	// Match [\x00-\u10ffff] lazily any number of times.
	lazyloop_pos1 = pos
	goto LazyLoopEnd1

LazyLoopBacktrack1:
	r.UncaptureUntil(lazyloop_capturepos1)
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	pos = lazyloop_pos1
	slice = r.Runtext[pos:]
	if len(slice) == 0 || false {
		goto CaptureBacktrack1
	}
	pos++
	slice = r.Runtext[pos:]
	lazyloop_pos1 = helpers.IndexOf(slice, []rune("[/"))
	if lazyloop_pos1 < 0 {
		goto CaptureBacktrack1
	}
	pos += lazyloop_pos1
	slice = r.Runtext[pos:]
	lazyloop_pos1 = pos

LazyLoopEnd1:
	lazyloop_capturepos1 = r.Crawlpos()

	r.Capture(3, capture_starting_pos2, pos)

	goto CaptureSkipBacktrack2

CaptureBacktrack2:
	goto LazyLoopBacktrack1

CaptureSkipBacktrack2:
	;

	// Node: Multi-M(String = "[/")
	// Match the string "[/".
	if !helpers.StartsWith(slice, []rune("[/")) {
		goto CaptureBacktrack2
	}

	// Node: Ref-M(index = 1)
	// Match the same text as matched by the "1" capture group.
	pos += 2
	slice = r.Runtext[pos:]

	// If the "1" capture group hasn't matched, the backreference doesn't match.
	if !r.IsMatched(1) {
		goto CaptureBacktrack2
	}

	// Get the captured text.  If it doesn't match at the current position, the backreference doesn't match.
	matchLength = r.MatchLength(1)
	if len(slice) < matchLength || !helpers.Equals(r.Runtext, r.MatchIndex(1), matchLength, slice[:matchLength]) {
		goto CaptureBacktrack2
	}
	pos += matchLength
	slice = r.Runtext[pos:]

	// Node: One-M(Ch = \])
	// Match ']'.
	if len(slice) == 0 || slice[0] != ']' {
		goto CaptureBacktrack2
	}

	// The input matched.
	pos++
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture-M(index = 0, unindex = -1)
 Concatenate-M
  One-M(Ch = \[)
  Capture-M(index = 1, unindex = -1)
   Concatenate-M
    Setloop-M(Set = [\w])(Min = 1, Max = inf)
    Boundary-M
  Capture-M(index = 2, unindex = -1)
   Notonelazy-M(Ch = \n)(Min = 0, Max = inf)
  One-M(Ch = \])
  Capture-M(index = 3, unindex = -1)
   Setlazy-M(Set = [\x00-\u10ffff])(Min = 0, Max = inf)
  Multi-M(String = "[/")
  Ref-M(index = 1)
  One-M(Ch = \])
*/
// From bb.go:39:32
// Pattern: "\\[(\\w+\\b)(.*?)\\]([\\s\\S]*?)\\[\\/\\1\\]"
// Options: regexp2.Multiline
type custom_Engine struct{}

func (custom_Engine) Caps() map[int]int        { return nil }
func (custom_Engine) CapNames() map[string]int { return nil }
func (custom_Engine) CapsList() []string       { return nil }
func (custom_Engine) CapSize() int             { return 4 }

func (custom_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 6 characters
	if pos <= len(r.Runtext)-6 {
		// The pattern begins with [\[]
		// Find the next occurrence. If it can't be found, there's no match.
		span := r.Runtext[pos:]
		for i := 0; i < len(span)-5; i++ {
			indexOfPos := helpers.IndexOfAny1(span[i:], '[')
			if indexOfPos < 0 {
				goto NoMatchFound
			}
			i += indexOfPos

			if helpers.IsWordChar(span[i+1]) {
				r.Runtextpos = pos + i
				return true
			}
		}
	}

	// No match found
NoMatchFound:
	r.Runtextpos = len(r.Runtext)
	return false
}

func (custom_Engine) Execute(r *regexp2.Runner) error {
	capture_starting_pos := 0
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	iteration := 0
	charloop_capture_pos := 0
	capture_starting_pos1 := 0
	lazyloop_capturepos := 0
	lazyloop_pos := 0
	capture_starting_pos2 := 0
	lazyloop_capturepos1 := 0
	lazyloop_pos1 := 0
	matchLength := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Concatenate-M
	// Node: One-M(Ch = \[)
	// Match '['.
	if len(slice) == 0 || slice[0] != '[' {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	// Node: Capture-M(index = 1, unindex = -1)
	// "1" capture group
	pos++
	slice = r.Runtext[pos:]
	capture_starting_pos = pos

	// Node: Concatenate-M
	// Node: Setloop-M(Set = [\w])(Min = 1, Max = inf)
	// Match [\w] greedily at least once.
	charloop_starting_pos = pos

	iteration = 0
	for iteration < len(slice) && helpers.IsWordChar(slice[iteration]) {
		iteration++
	}

	if iteration == 0 {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}

	slice = slice[iteration:]
	pos += iteration

	charloop_ending_pos = pos
	charloop_starting_pos++
	goto CharLoopEnd

CharLoopBacktrack:
	r.UncaptureUntil(charloop_capture_pos)

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}
	charloop_ending_pos--
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:
	charloop_capture_pos = r.Crawlpos()

	// Node: Boundary-M
	// Match if at a word boundary.
	if !r.IsBoundary(pos) {
		goto CharLoopBacktrack
	}

	r.Capture(1, capture_starting_pos, pos)

	goto CaptureSkipBacktrack

CaptureBacktrack:
	goto CharLoopBacktrack

CaptureSkipBacktrack:
	;

	// Node: Capture-M(index = 2, unindex = -1)
	// "2" capture group
	capture_starting_pos1 = pos

	// Node: Notonelazy-M(Ch = \n)(Min = 0, Max = inf)
	// Match a character other than '\n' lazily any number of times.
	lazyloop_pos = pos
	goto LazyLoopEnd

LazyLoopBacktrack:
	r.UncaptureUntil(lazyloop_capturepos)
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	pos = lazyloop_pos
	slice = r.Runtext[pos:]
	if len(slice) == 0 || slice[0] == '\n' {
		goto CaptureBacktrack
	}
	pos++
	slice = r.Runtext[pos:]
	lazyloop_pos = helpers.IndexOfAny2(slice, '\n', ']')
	if lazyloop_pos >= len(slice) || slice[lazyloop_pos] == '\n' {
		goto CaptureBacktrack
	}
	pos += lazyloop_pos
	slice = r.Runtext[pos:]
	lazyloop_pos = pos

LazyLoopEnd:
	lazyloop_capturepos = r.Crawlpos()

	r.Capture(2, capture_starting_pos1, pos)

	goto CaptureSkipBacktrack1

CaptureBacktrack1:
	goto LazyLoopBacktrack

CaptureSkipBacktrack1:
	;

	// Node: One-M(Ch = \])
	// Match ']'.
	if len(slice) == 0 || slice[0] != ']' {
		goto CaptureBacktrack1
	}

	// Node: Capture-M(index = 3, unindex = -1)
	// "3" capture group
	pos++
	slice = r.Runtext[pos:]
	capture_starting_pos2 = pos

	// Node: Setlazy-M(Set = [\x00-\u10ffff])(Min = 0, Max = inf)
	// Match [\x00-\u10ffff] lazily any number of times.
	lazyloop_pos1 = pos
	goto LazyLoopEnd1

LazyLoopBacktrack1:
	r.UncaptureUntil(lazyloop_capturepos1)
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	pos = lazyloop_pos1
	slice = r.Runtext[pos:]
	if len(slice) == 0 || false {
		goto CaptureBacktrack1
	}
	pos++
	slice = r.Runtext[pos:]
	lazyloop_pos1 = helpers.IndexOf(slice, []rune("[/"))
	if lazyloop_pos1 < 0 {
		goto CaptureBacktrack1
	}
	pos += lazyloop_pos1
	slice = r.Runtext[pos:]
	lazyloop_pos1 = pos

LazyLoopEnd1:
	lazyloop_capturepos1 = r.Crawlpos()

	r.Capture(3, capture_starting_pos2, pos)

	goto CaptureSkipBacktrack2

CaptureBacktrack2:
	goto LazyLoopBacktrack1

CaptureSkipBacktrack2:
	;

	// Node: Multi-M(String = "[/")
	// Match the string "[/".
	if !helpers.StartsWith(slice, []rune("[/")) {
		goto CaptureBacktrack2
	}

	// Node: Ref-M(index = 1)
	// Match the same text as matched by the "1" capture group.
	pos += 2
	slice = r.Runtext[pos:]

	// If the "1" capture group hasn't matched, the backreference doesn't match.
	if !r.IsMatched(1) {
		goto CaptureBacktrack2
	}

	// Get the captured text.  If it doesn't match at the current position, the backreference doesn't match.
	matchLength = r.MatchLength(1)
	if len(slice) < matchLength || !helpers.Equals(r.Runtext, r.MatchIndex(1), matchLength, slice[:matchLength]) {
		goto CaptureBacktrack2
	}
	pos += matchLength
	slice = r.Runtext[pos:]

	// Node: One-M(Ch = \])
	// Match ']'.
	if len(slice) == 0 || slice[0] != ']' {
		goto CaptureBacktrack2
	}

	// The input matched.
	pos++
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

func init() {
	regexp2.RegisterEngine("```([\\s\\S])```", regexp2.Multiline, &largeCodeConv_Engine{})
	regexp2.RegisterEngine("\\[code\\]([\\s\\S]*?)\\[\\/code\\]", regexp2.Multiline, &code_Engine{})
	regexp2.RegisterEngine("\\[(b|bold|i|italics|u|underline|s|strike|font|size|color|smallcaps|url|link|youtube|img|image|title|t[1-6]|align|float|ul|bullet|ol|number)(.*?)\\]([\\s\\S]*?)\\[\\/\\1\\]", regexp2.Multiline, &main_Engine{})
	regexp2.RegisterEngine("\\[(\\w+\\b)(.*?)\\]([\\s\\S]*?)\\[\\/\\1\\]", regexp2.Multiline, &custom_Engine{})
	var _ = helpers.Min
	var _ = syntax.NewCharSetRuntime
	var _ = unicode.IsDigit
}
