package util

import "regexp"

var specialCharRegex *regexp.Regexp
var specialCharWithSpaceRegex *regexp.Regexp
var allCharsButAlphabetRegex *regexp.Regexp

func init() {
	reg := regexp.MustCompile("[^a-zA-Z0-9äöüÄÖÜß]+")
	specialCharRegex = reg
	reg = regexp.MustCompile("[^a-zA-Z0-9äöüÄÖÜß ]+")
	specialCharWithSpaceRegex = reg
	reg = regexp.MustCompile("[^a-zA-ZäöüÄÖÜß]+")
	allCharsButAlphabetRegex = reg
}
