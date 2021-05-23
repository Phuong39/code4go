/**
* 作者：刘时明
* 时间：2019/8/29-19:00
* 作用：基于Go的脚本语言
 */
package main

import (
	. "courage/compiler/lexer"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// crgc hello.c
	path := "resources/hello.c"
	if data, err := ioutil.ReadFile(path); err == nil {
		parseSource(string(data), getSourceName(&path))
	} else {
		panic(err)
	}
}

func getSourceName(p *string) string {
	end := strings.LastIndex(*p, ".")
	if end < 0 {
		panic("不支持的路径")
	}
	start := strings.LastIndex(*p, "/")
	if start < 0 {
		start = -1
	}
	start++
	return (*p)[start:end]
}

func parseSource(chunk string, chunkName string) {
	lexer := NewLexer(chunk, chunkName)
	for {
		line, kind, token := lexer.NextToken()
		fmt.Printf("行号=[%2d] 类型=[%-10s] token > %s\n", line, kindToCategory(kind), token)
		if kind == TOKEN_EOF {
			break
		}
	}
}

func kindToCategory(kind int) string {
	switch {
	case kind < TOKEN_SEP_SEMI:
		return "other"
	case kind <= TOKEN_SEP_RCURLY:
		return "separator"
	case kind <= TOKEN_OP_NOT:
		return "operator"
	case kind <= TOKEN_KW_WHILE:
		return "keyword"
	case kind == TOKEN_IDENTIFIER:
		return "identifier"
	case kind == TOKEN_NUMBER:
		return "number"
	case kind == TOKEN_STRING:
		return "string"
	default:
		return "other"
	}
}
