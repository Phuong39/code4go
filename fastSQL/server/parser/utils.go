package parser

// 是否字母
func isLetter(c rune) bool {
	return (c >= 97 && c <= 122) || (c >= 65 && c <= 90)
}

// 是否数字
func isDigit(c rune) bool {
	return c >= 48 && c <= 57
}

// 是否字母或数字
func isLetterOrDigit(c rune) bool {
	return isLetter(c) || isDigit(c)
}

// 是否空格
func isSpace(c rune) bool {
	return c == ' '
}

// 是否逗号
func isComma(c rune) bool {
	return c == ','
}
