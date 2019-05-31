package main

import (
	"fmt"
	"math/rand"
	"time"
)

func a(a int, b ...int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%d\n%d\n", 3, rand.Intn(10))
}

func swap(x, y string) (a string, b string) {
	return y, x
}

func main() {

	fmt.Printf("hello world\n")
	fmt.Println(swap("a", "b"))
	a(3, 2)

}

/* FunctionDecl = "func" FunctionName Signature [FunctionBody] . */
/* Signature = Parametes [Result] */
/* Result =  Parametes | Type */
/* Parameters = "(" [ParameterList [,]] ")"  */
/* ParametersList = ParameterDecl {"," ParameterDecl } */
/* ParameterDecl = [IdentifierList] [ "..." ] Type */
/* IdentifierList = Identifier {"," Identifier} */
/* Identifier = letter {letter | unicode_digit } */
