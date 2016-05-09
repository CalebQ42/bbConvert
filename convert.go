package bbConvert

import "strings"

func bbconv(input string) string {
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			for j := i; j < len(input); j++ {
				if input[j] == ']' && checktags(input[i:j+1]) {
					input = input[:i] + convert(input[i:j+1]) + input[j+1:]
					i = 0
					break
				}
			}
		}
	}
	return input
}

func convert(input string) string {
	var front, meat string
	for i, v := range input {
		if v == ']' {
			front = input[1:i]
			for j := len(input) - 1; j > i; j-- {
				if input[j] == '[' {
					meat = input[i+1 : j]
					break
				}
			}
			break
		}
	}
	out := toHTML(front, bbconv(meat))
	return out
}

func checktags(input string) bool {
	input = strings.ToLower(input)
	c := make(chan string, 2)
	go checkfront(input, c)
	go checkback(input, c)
	fr, bk := <-c, <-c
	if fr == "Nope" || bk == "Nope" {
		return false
	}
	if fr == bk {
		return true
	}
	return false
}

func checkfront(input string, channel chan string) {
	for i, v := range input {
		if v == ' ' || v == '=' || v == ']' {
			channel <- input[1:i]
			return
		}
	}
	channel <- "Nope"
}

func checkback(input string, channel chan string) {
	for i := len(input) - 1; i > -1; i-- {
		v := input[i]
		if v == '[' {
			channel <- input[i+2 : len(input)-1]
			return
		}
	}
	channel <- "Nope"
}

func indexAll(s, set string) []int {
	indexi := make([]int, strings.Count(s, set))
	orig := s
	for i := range indexi {
		if i > 0 {
			indexi[i] = strings.Index(s, set) + indexi[i-1] + 1
		} else {
			indexi[i] = strings.Index(s, set)
		}
		s = orig[indexi[i]+1:]
	}
	return indexi
}
