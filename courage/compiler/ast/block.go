/*
* 作者：刘时明
* 时间：2020/5/26-22:35
* 作用：
 */
package ast

type Block struct {
	LastLine int
	Stats    []Stat
	RetExps  []Exp
}
