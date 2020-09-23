/*
* 作者：刘时明
* 时间：2020/5/26-22:36
* 作用：
 */
package ast

type Stat interface{}

type EmptyStat struct{}

type BreakStat struct {
	Line int
}

type LabelStat struct {
	Name string
}

type GotoStat struct {
	Name string
}

type DoStat struct {
	Block *Block
}

type FuncCallStat = FuncallExp

type WhileStat struct {
	Exp   Exp
	Block *Block
}

type RepeatStat struct {
	Exp   Exp
	Block *Block
}

type IfStat struct {
	Exp   []Exp
	Block []*Block
}

type ForNumStat struct {
	LineOfFor int
	LineOfDo  int
	VarName   string
	InitExp   Exp
	LimitExp  Exp
	StepExp   Exp
	Block     *Block
}

type ForInStat struct {
	LineOfFor int
	NameList  []string
	ExpList   []Exp
	Block     *Block
}

type LocalVarDeclStat struct {
	LastLine int
	NameList []string
	ExpList  []Exp
}

type AssignStat struct {
	LastLine int
	VarList  []Exp
	ExpList  []Exp
}

type LocalFuncDefStat struct {
	Name string
	Exp  *FuncDefExp
}
