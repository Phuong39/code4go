package parser

import (
	"bytes"
	"strings"
)

func Tokenizer(sql string) []string {
	if !strings.HasSuffix(sql, ";") {
		sql = sql + ";"
	}
	list := make([]string, 0, 50)
	var current int
	var c rune
	var buf bytes.Buffer
	for current < len(sql) {
		c = rune(sql[current])
		// 是否字母
		if isLetter(c) {
			buf.Reset()
			for isLetterOrDigit(c) {
				buf.WriteRune(c)
				current++
				c = rune(sql[current])
			}
			if c == '.' {
				buf.WriteRune(c)
				current++
				c = rune(sql[current])
				if c == '*' {
					buf.WriteRune(c)
					current++
					c = rune(sql[current])
				} else {
					for isLetterOrDigit(c) {
						buf.WriteRune(c)
						current++
						c = rune(sql[current])
					}
				}
			}
			// 转为大写字母
			list = append(list, strings.ToUpper(buf.String()))
			continue
		}

		// 是否数字
		if isDigit(c) {
			buf.Reset()
			for isDigit(c) {
				buf.WriteRune(c)
				current++
				c = rune(sql[current])
			}
			list = append(list, buf.String())
			continue
		}

		// 是否空格或者逗号
		if isSpace(c) || isComma(c) {
			current++
			continue
		}

		if c == ';' {
			break
		}

		switch c {
		case '(':
			fallthrough
		case '[':
			list = append(list, string(c))
			current++
		case ')':
			fallthrough
		case ']':
			list = append(list, string(c))
			current++
		case '\'':
			buf.Reset()
			buf.WriteString("'")
			//设置last以判断是否是转义的'
			lastC := c
			for {
				current++
				c = rune(sql[current])
				if c == '\'' && lastC != '\\' {
					buf.WriteRune(c)
					list = append(list, buf.String())
					break
				}
				buf.WriteRune(c)
				lastC = c
			}
			current++
		case '=':
			list = append(list, string(c))
			current++
		case '>':
			if sql[current+1] == '=' {
				list = append(list, ">=")
				current += 2
			} else {
				list = append(list, ">")
				current++
			}
		case '<':
			if sql[current+1] == '=' {
				list = append(list, "<=")
				current += 2
			} else {
				list = append(list, "<")
				current++
			}
		case '!':
			if sql[current+1] != '=' {
				panic("Syntax error, '!' must be before '='.")
			}
			list = append(list, "!=")
			current += 2
			break
		case ',':
			list = append(list, string(c))
			current++
			break
		case '*':
			list = append(list, string(c))
			current++
		default:
			panic("Unrecognized characters.")
		}
	}
	return list
}
