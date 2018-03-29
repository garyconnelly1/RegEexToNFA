
package main
//Thompsons algorithm
import (
	"fmt"
	
)

//create structers
type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}

type nfa struct{
	initial *state
	accept *state
}

//function that returns a pointer to the nfa struct
func poregtonfa(pofix string) *nfa{
	//create an array that represents the nfa stack
	nfastack := []*nfa{}

//for each element in the string that was passed in to the function
	for _, r := range pofix {
		switch r{
			//if the character = '.'
		case '.':
			//variable frag 2 is assigned the value of the last element of the nfastack array
			frag2 := nfastack[len(nfastack)-1]
			//nfa stack is set to up to the value of the last element in the nfa stack array
			nfastack = nfastack[:len(nfastack)-1]
			//variable frag 1 is assigned the value of the last element of the nfastack array
			frag1 := nfastack[len(nfastack)-1]
			//nfa stack is set to up to the value of the last element in the nfa stack array
			nfastack = nfastack[:len(nfastack)-1]

			//edge 1, which is pointed to by the accept in the nfa struct, which pointed to by the nfa stack, is set to the initial state
			//of frag2.
			frag1.accept.edge1 = frag2.initial

			//append the nfa with the initial state of frag1 and the accept state of frag 2 to the nfa stack.
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

			//if the character = '|'
		case '|':
			//variable frag 2 is assigned the value of the last element of the nfastack array
			frag2 := nfastack[len(nfastack)-1]
			//nfa stack is set to up to the value of the last element in the nfa stack array
			nfastack = nfastack[:len(nfastack)-1]
			//variable frag 1 is assigned the value of the last element of the nfastack array
			frag1 := nfastack[len(nfastack)-1]
			//nfa stack is set to up to the value of the last element in the nfa stack array
			nfastack = nfastack[:len(nfastack)-1]

			//generate the new states states
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			//frag 1 and 2, on edge one of accept, is given accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//append these states to the nfa stack array
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			//if the character = '*'
		case '*':
			//variable frag is assigned the value of the last element of the nfastack array
			frag := nfastack[len(nfastack)-1]
			//nfa stack is set to up to the value of the last element in the nfa stack array
			nfastack = nfastack[:len(nfastack)-1]

			//accept is initialised to a new state
			accept := state{}
			//initial is initialised to initial state of frag and accept state of edge 2
			initial := state{edge1: frag.initial, edge2: &accept}
			// set the frag accept states
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//append initial and accept statse to the nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		}//end switch
	}//end for range

	return nfastack[0]
}//end function

//add state function
func addState(l []*state, s *state, a *state) []*state {
		return l
}//end addState


func main(){

	fmt.Println(pomatch("ab.*c*|", "abab"))
}
