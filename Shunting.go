//shunting yard algorithm

package main

import (
	"fmt"
)

//function for changing infix to post fix notation
func intopost(infix string)string{
	//initialize variables needed
	specials := map[rune]int{'*':10, '.': 9, '|':8}
	pofix := []rune{}
	s := []rune{} 
	return infix
}

//main function
func main(){
	// postfix = ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	// postfix = abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))
	
	// postfix = abb.+.c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))

	// postfix = abb.+.c.
	fmt.Println("Infix: ", "n.a.m.e")
	fmt.Println("Postfix: ", intopost("n.a.m.e"))
}