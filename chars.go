package util

func CutSpecialChars(str string) string {
	return specialCharRegex.ReplaceAllString(str, "")
}
func CutSpecialCharsKeepSpace(str string) string {
	return specialCharWithSpaceRegex.ReplaceAllString(str, "")
}
func CutAllCharsButAlphabet(str string) string {
	return allCharsButAlphabetRegex.ReplaceAllString(str, "")
}
