package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func random(a int, b ...int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%d\n%d\n", 3, rand.Intn(10))
}

func swap(x, y string) (a string, b string) {
	return y, x
}

func main() {

	fmt.Printf("hello world\n")
	fmt.Println(swap("a", "b"))
	random(2, 3, 4, 5)
	var x, y int = 3, 2
	var z = math.Sqrt(float64(x*x + y*y))
	var t = uint(2)
	var i = 0.2 + 0.5i
	const (
		a = iota
		b
		c
	)
	const (
		n  = 3
		n2 = iota
		n3
	)
	fmt.Println(string(x), y, z, t, i, a, b, 2222, n, n2, n3, a, b, c)
}

/* FunctionDecl = "func" FunctionName Signature [FunctionBody] . */
/* Signature = Parametes [Result] */
/* Result =  Parametes | Type */
/* Parameters = "(" [ParameterList [,]] ")"  */
/* ParametersList = ParameterDecl {"," ParameterDecl } */
/* ParameterDecl = [IdentifierList] [ "..." ] Type */
/* IdentifierList = Identifier {"," Identifier} */
/* Identifier = letter {letter | unicode_digit } */

/* InterfaceType = "interface" "{" { MethodSpec";"} "}" */
/* MethodSpec = MethodName Signature | InterfaceName */

/* VarDecl = "var" (VarSpec | "("{ VarSpec [;]} ")") */
/* VarSpec = IdentifeierList (Type ["=" ExpressionList] | "=" ExpressionList )*/
/* ConstSpec = IdentifierList [[Type] "=" ExpressionList] */

func varDecl() {
	var (
		abcd, bcde, cdef int
		defg             = ""
	)
	println(abcd, bcde, cdef, defg)
}

/* ForStmt = "for" [Condition | ForClause | RangeClause] Block */
/* Condition = Expression */
/* ForClause = [InitStmt] ";" [Condition]";" [PostStmt]  */
/* InitStmt = SimpleStmt */
/* PostStmt = SimpleStmt
/* RangeClause = [ ExpressionList "=" | IdentifierList ":="] "range" Expression */

func funcname(a, b, c int) (d, e int) {
	for 1 < 2 {

	}
	for i := 0; 2 < 3; i++ {

	}
	words := []string{"a", "b", "c"}
	var f string
	for _, f = range words {
		fmt.Println(f)
	}
	return 3, 5
}

/* IfStmt = "if" [SimpleStmt ";"] Expression Block ["else" (IfStmt | Block) ] */
func ift() {
	if x := 3; x < 4 {

	}
}

/* SwitchStmt = "switch" [ SimpleStmt ";"] [Expression] "{" {ExprCaseClause} "}" */
/* ExprCaseClause = ExpreSwitchCase ":" StatementList */
/* ExprSwitchCase = "case" ExpressionList | "default" */
func switchstmt() {
	switch x := 2; x {
	case 2, 5, 8:
		println(x)
	case 3:
		println(x)
	default:
		println(x)
	}
	switch {
	case time.Now().Hour() < 12:
		println(3)
	default:
		println("default")
	}
}

/* DeferStmt = "defer" Expression */
func deferstmt() {
	defer println("defferred")
	println("start")
}

/* Go Basic Types */
/*

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

*/
