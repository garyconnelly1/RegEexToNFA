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

//for each character in the expression
	for _, r := range infix{
		switch {
			
		case r == '(' :
			s = append(s,r)

		case r == ')' :
			for(s[len(s)-1] != '('){
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]  
			}//end for
			s = s[:len(s)-1]  

		case specials[r] > 0 :
			//for
			for(len(s) > 0 && specials[r] <= specials[s[len(s)-1]] ){
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] 
			}//end for
			s = append(s,r)
		default:
		pofix = append(pofix,r)
		}
	
	}//end for range


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