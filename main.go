package bbConvert

import "regexp"

func ConvertBBCode(in string) string {
	bbMatch, err := regexp.Compile(BBMatchRegEx)
	if err != nil {
		return ""
	}
	ind := bbMatch.FindStringIndex(in)
	for ind != nil {

	}
}

func convertMatchedTags(tags string) string {
	return tags
}
