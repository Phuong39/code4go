package parser

import . "courage/compiler/ast"
import . "courage/compiler/lexer"

func parsePrefixExp(lexer *Lexer) Exp {
	var exp Exp
	if lexer.LookAhead() == TOKEN_IDENTIFIER {
		line, name := lexer.NextIdentifier()
		exp = &NameExp{Line: line, Name: name}
	} else { // ‘(’ exp ‘)’
		exp = parseParensExp(lexer)
	}
	return _finishPrefixExp(lexer, exp)
}

func parseParensExp(lexer *Lexer) Exp {
	lexer.NextTokenOfKind(TOKEN_SEP_LPAREN) // (
	exp := parseExp(lexer)                  // exp
	lexer.NextTokenOfKind(TOKEN_SEP_RPAREN) // )
	switch exp.(type) {
	case *VarargExp, *FuncCallExp, *NameExp, *TableAccessExp:
		return &ParensExp{Exp: exp}
	}
	// no need to keep parens
	return exp
}

func _finishPrefixExp(lexer *Lexer, exp Exp) Exp {
	for {
		switch lexer.LookAhead() {
		case TOKEN_SEP_LBRACK: // prefixexp ‘[’ exp ‘]’
			lexer.NextToken()                       // ‘[’
			keyExp := parseExp(lexer)               // exp
			lexer.NextTokenOfKind(TOKEN_SEP_RBRACK) // ‘]’
			exp = &TableAccessExp{LastLine: lexer.Line(), PrefixExp: exp, KeyExp: keyExp}
		case TOKEN_SEP_DOT: // prefixexp ‘.’ Name
			lexer.NextToken()                    // ‘.’
			line, name := lexer.NextIdentifier() // Name
			keyExp := &StringExp{Line: line, Str: name}
			exp = &TableAccessExp{LastLine: line, PrefixExp: exp, KeyExp: keyExp}
		case TOKEN_SEP_COLON, // prefixexp ‘:’ Name args
			TOKEN_SEP_LPAREN, TOKEN_SEP_LCURLY, TOKEN_STRING: // prefixexp args
			exp = _finishFuncCallExp(lexer, exp)
		default:
			return exp
		}
	}
}

func _finishFuncCallExp(lexer *Lexer, prefixExp Exp) *FuncCallExp {
	nameExp := _parseNameExp(lexer)
	line := lexer.Line()
	args := _parseArgs(lexer)
	lastLine := lexer.Line()
	return &FuncCallExp{Line: line, LastLine: lastLine, PrefixExp: prefixExp, NameExp: nameExp, Args: args}
}

func _parseNameExp(lexer *Lexer) *StringExp {
	if lexer.LookAhead() == TOKEN_SEP_COLON {
		lexer.NextToken()
		line, name := lexer.NextIdentifier()
		return &StringExp{Line: line, Str: name}
	}
	return nil
}

func _parseArgs(lexer *Lexer) (args []Exp) {
	switch lexer.LookAhead() {
	case TOKEN_SEP_LPAREN:
		lexer.NextToken()
		if lexer.LookAhead() != TOKEN_SEP_RPAREN {
			args = parseExpList(lexer)
		}
		lexer.NextTokenOfKind(TOKEN_SEP_RPAREN)
	case TOKEN_SEP_LCURLY:
		args = []Exp{parseTableConstructorExp(lexer)}
	default:
		line, str := lexer.NextTokenOfKind(TOKEN_STRING)
		args = []Exp{&StringExp{Line: line, Str: str}}
	}
	return
}
