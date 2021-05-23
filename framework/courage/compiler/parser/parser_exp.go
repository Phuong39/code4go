/*
* 作者：刘时明
* 时间：2020/5/27-0:08
* 作用：
 */
package parser

import (
	. "courage/compiler/ast"
	. "courage/compiler/lexer"
	"courage/number"
)

func parseExp(lexer *Lexer) Exp {
	return parseExp12(lexer)
}

func parseExpList(lexer *Lexer) []Exp {
	exps := make([]Exp, 0, 4)
	exps = append(exps, parseExp(lexer))
	for lexer.LookAhead() == TOKEN_SEP_COMMA {
		lexer.NextToken()
		exps = append(exps, parseExp(lexer))
	}
	return exps
}

func parseFuncDefExp(lexer *Lexer) *FuncDefExp {
	line := lexer.Line()                               // function
	lexer.NextTokenOfKind(TOKEN_SEP_LPAREN)            // (
	parList, isVararg := _parseParList(lexer)          // [parlist]
	lexer.NextTokenOfKind(TOKEN_SEP_RPAREN)            // )
	block := parseBlock(lexer)                         // block
	lastLine, _ := lexer.NextTokenOfKind(TOKEN_KW_END) // end
	return &FuncDefExp{Line: line, LastLine: lastLine, ParList: parList, IsVararg: isVararg, Block: block}
}

func _parseParList(lexer *Lexer) (names []string, isVararg bool) {
	switch lexer.LookAhead() {
	case TOKEN_SEP_RPAREN:
		return nil, false
	case TOKEN_VARARG:
		lexer.NextToken()
		return nil, true
	}
	_, name := lexer.NextIdentifier()
	names = append(names, name)
	for lexer.LookAhead() == TOKEN_SEP_COMMA {
		lexer.NextToken()
		if lexer.LookAhead() == TOKEN_IDENTIFIER {
			_, name := lexer.NextIdentifier()
			names = append(names, name)
		} else {
			lexer.NextTokenOfKind(TOKEN_VARARG)
			isVararg = true
			break
		}
	}
	return
}

func parseTableConstructorExp(lexer *Lexer) *TableConstructorExp {
	line := lexer.Line()
	lexer.NextTokenOfKind(TOKEN_SEP_LCURLY)
	keyExps, valExps := _parseFieldList(lexer)
	lexer.NextTokenOfKind(TOKEN_SEP_RCURLY)
	lastLine := lexer.Line()
	return &TableConstructorExp{Line: line, LastLine: lastLine, KeyExps: keyExps, ValExps: valExps}
}

func _parseFieldList(lexer *Lexer) (ks, vs []Exp) {
	if lexer.LookAhead() != TOKEN_SEP_RCURLY {
		k, v := _parseField(lexer)
		ks = append(ks, k)
		vs = append(vs, v)
		for _isFieldSep(lexer.LookAhead()) {
			lexer.NextToken()
			if lexer.LookAhead() != TOKEN_SEP_RCURLY {
				k, v := _parseField(lexer)
				ks = append(ks, k)
				vs = append(vs, v)
			} else {
				break
			}
		}
	}
	return
}

func _isFieldSep(tokenKind int) bool {
	return tokenKind == TOKEN_SEP_COMMA || tokenKind == TOKEN_SEP_SEMI
}

func _parseField(lexer *Lexer) (k, v Exp) {
	if lexer.LookAhead() == TOKEN_SEP_LBRACK {
		lexer.NextToken()                       // [
		k = parseExp(lexer)                     // exp
		lexer.NextTokenOfKind(TOKEN_SEP_RBRACK) // ]
		lexer.NextTokenOfKind(TOKEN_OP_ASSIGN)  // =
		v = parseExp(lexer)                     // exp
		return
	}
	exp := parseExp(lexer)
	if nameExp, ok := exp.(*NameExp); ok {
		if lexer.LookAhead() == TOKEN_OP_ASSIGN {
			// Name ‘=’ exp => ‘[’ LiteralString ‘]’ = exp
			lexer.NextToken()
			k = &StringExp{Line: nameExp.Line, Str: nameExp.Name}
			v = parseExp(lexer)
			return
		}
	}
	return nil, exp
}

func parseNumberExp(lexer *Lexer) Exp {
	line, _, token := lexer.NextToken()
	if i, ok := number.ParseInteger(token); ok {
		return &IntegerExp{line, i}
	} else if f, ok := number.ParseFloat(token); ok {
		return &FloatExp{line, f}
	} else { // todo
		panic("not a number: " + token)
	}
}

func parseExp0(lexer *Lexer) Exp {
	switch lexer.LookAhead() {
	case TOKEN_VARARG: // ...
		line, _, _ := lexer.NextToken()
		return &VarargExp{Line: line}
	case TOKEN_KW_NULL: // nil
		line, _, _ := lexer.NextToken()
		return &NullExp{Line: line}
	case TOKEN_KW_TRUE: // true
		line, _, _ := lexer.NextToken()
		return &TrueExp{Line: line}
	case TOKEN_KW_FALSE: // false
		line, _, _ := lexer.NextToken()
		return &FalseExp{Line: line}
	case TOKEN_STRING: // LiteralString
		line, _, token := lexer.NextToken()
		return &StringExp{Line: line, Str: token}
	case TOKEN_NUMBER: // Numeral
		return parseNumberExp(lexer)
	case TOKEN_SEP_LCURLY: // tableconstructor
		return parseTableConstructorExp(lexer)
	case TOKEN_KW_FUNCTION: // functiondef
		lexer.NextToken()
		return parseFuncDefExp(lexer)
	default: // prefixexp
		return parsePrefixExp(lexer)
	}
}

func parseExp1(lexer *Lexer) Exp { // pow is right associative
	exp := parseExp0(lexer)
	if lexer.LookAhead() == TOKEN_OP_POW {
		line, op, _ := lexer.NextToken()
		exp = &BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp2(lexer)}
	}
	return optimizePow(exp)
}

func parseExp2(lexer *Lexer) Exp {
	switch lexer.LookAhead() {
	case TOKEN_OP_UNM, TOKEN_OP_BNOT, TOKEN_OP_LEN, TOKEN_OP_NOT:
		line, op, _ := lexer.NextToken()
		exp := &UnopExp{Line: line, Op: op, Exp: parseExp2(lexer)}
		return optimizeUnaryOp(exp)
	}
	return parseExp1(lexer)
}

func parseExp3(lexer *Lexer) Exp {
	exp := parseExp2(lexer)
	for {
		switch lexer.LookAhead() {
		case TOKEN_OP_MUL, TOKEN_OP_MOD, TOKEN_OP_DIV, TOKEN_OP_IDIV:
			line, op, _ := lexer.NextToken()
			exp = optimizeArithBinaryOp(&BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp2(lexer)})
		default:
			return exp
		}
	}
}

func parseExp4(lexer *Lexer) Exp {
	exp := parseExp3(lexer)
	for {
		switch lexer.LookAhead() {
		case TOKEN_OP_ADD, TOKEN_OP_SUB:
			line, op, _ := lexer.NextToken()
			exp = optimizeArithBinaryOp(&BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp3(lexer)})
		default:
			return exp
		}
	}
}

func parseExp5(lexer *Lexer) Exp {
	exp := parseExp4(lexer)
	if lexer.LookAhead() != TOKEN_OP_CONCAT {
		return exp
	}
	line := 0
	exps := []Exp{exp}
	for lexer.LookAhead() == TOKEN_OP_CONCAT {
		line, _, _ = lexer.NextToken()
		exps = append(exps, parseExp4(lexer))
	}
	return &ConcatExp{Line: line, Exps: exps}
}

func parseExp6(lexer *Lexer) Exp {
	exp := parseExp5(lexer)
	for {
		switch lexer.LookAhead() {
		case TOKEN_OP_SHL, TOKEN_OP_SHR:
			line, op, _ := lexer.NextToken()
			exp = optimizeBitwiseBinaryOp(&BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp5(lexer)})
		default:
			return exp
		}
	}
}

func parseExp7(lexer *Lexer) Exp {
	exp := parseExp6(lexer)
	for lexer.LookAhead() == TOKEN_OP_BAND {
		line, op, _ := lexer.NextToken()
		exp = optimizeBitwiseBinaryOp(&BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp6(lexer)})
	}
	return exp
}

func parseExp8(lexer *Lexer) Exp {
	exp := parseExp7(lexer)
	for lexer.LookAhead() == TOKEN_OP_BXOR {
		line, op, _ := lexer.NextToken()
		exp = optimizeBitwiseBinaryOp(&BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp7(lexer)})
	}
	return exp
}

func parseExp9(lexer *Lexer) Exp {
	exp := parseExp8(lexer)
	for lexer.LookAhead() == TOKEN_OP_BOR {
		line, op, _ := lexer.NextToken()
		exp = optimizeBitwiseBinaryOp(&BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp8(lexer)})
	}
	return exp
}

func parseExp10(lexer *Lexer) Exp {
	exp := parseExp9(lexer)
	for {
		switch lexer.LookAhead() {
		case TOKEN_OP_LT, TOKEN_OP_GT, TOKEN_OP_NE,
			TOKEN_OP_LE, TOKEN_OP_GE, TOKEN_OP_EQ:
			line, op, _ := lexer.NextToken()
			exp = &BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp9(lexer)}
		default:
			return exp
		}
	}
}

func parseExp11(lexer *Lexer) Exp {
	exp := parseExp10(lexer)
	for lexer.LookAhead() == TOKEN_OP_AND {
		line, op, _ := lexer.NextToken()
		land := &BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp10(lexer)}
		exp = optimizeLogicalAnd(land)
	}
	return exp
}

func parseExp12(lexer *Lexer) Exp {
	exp := parseExp11(lexer)
	for lexer.LookAhead() == TOKEN_OP_OR {
		line, op, _ := lexer.NextToken()
		lor := &BinopExp{Line: line, Op: op, Exp1: exp, Exp2: parseExp11(lexer)}
		exp = optimizeLogicalOr(lor)
	}
	return exp
}
