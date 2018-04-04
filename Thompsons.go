
package main
//Thompsons algorithm
import (
	"fmt"
	"bufio"
	"os"
	shunting "./ShuntingPackage"
	"io/ioutil"
	"strings"
	"strconv"
	
	
	
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
		case '+':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = &initial

			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		}//end switch
	}//end for range

	return nfastack[0]
}//end function

//add state function
//function returns an array of pointers to state struct
func addState(l []*state, s *state, a *state) []*state {
	//append the state pointer to the l array
	l = append(l,s)

	//if state is not = to a(accept) and the value of symbol(rune) is 0
	if s != a && s.symbol == 0 {
		//recursively append states edge1 and accept
		l = addState(l, s.edge1, a)
		//if states edge2 is not null(0) 
		if s.edge2 != nil {
			//and states edge 2 and accept
			l = addState(l, s.edge2, a)
		}
	} 
	//return array of pointers l 
		return l
}//end addState

//regex matching function
func pomatch(po string, s string) bool{
	//initialise instance variables
	ismatch := false

	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

//for range of s
//for each character r in string s
	for _, r:= range s{
		//for each var c in current array
		for _, c:= range current{
			//if symbol = current character r
			if c.symbol == r{
				//array next is set to array of pointers returned from addState function
				next = addState(next[:], c.edge1, ponfa.accept)
			}//end if
		}//end range current for
		current, next = next, []*state{}
	}//end range s for


	//for range of current
	// for each var c in currentarray
	for _, c:= range current{
		//if c == nfa accept state
		if c == ponfa.accept {
			//then it is a match
			ismatch = true
			//break out of loop
			break
		}//end if
		
	}//end range current for


	//return true/false result
	return ismatch
}












//function to trim the last two ascii characters off the end of the string
func TrimFix(s string) string{
	if len(s) > 0{
		s = s[:len(s)-2]
	}
	return s
}


func main(){

	count := 0
	i := 0

	//shunting := new Shunting()

	//get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Regular expression: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Before trim suffix" + text)
	//trim last 2 ascii characters from the end
	text = TrimFix(text)
	//convert the regular expression from infix to post fix
	fmt.Println("intopost + " + shunting.Intopost(text))

	expression := shunting.Intopost(text)

	//read in Gutenberg text file
	 b, err := ioutil.ReadFile("TextFile.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

 str := string(b)

 //split the string into token words
  s := strings.Split(str, " ")

   for _, word := range s{
		 i++
		// fmt.Println(pomatch(text, word))
		 if(pomatch(expression, word) == true){
			 count++
			fmt.Println("YUUUUUURRRRRRRRRTTTTT" +  " " + word + " appears " + strconv.Itoa(i) + " words in ")
		 }else{
			// fmt.Println("nahh")
		 }
	 }

	 if (count > 0) {
		
		 fmt.Println("Yes! The expression " + text + " exists in the text tile " + strconv.Itoa(count) + " times.")
	 } else {
		 fmt.Println("No! The expression " + text + " does not exist in the text document.")
	 }
	 fmt.Println(pomatch("I+", "I"))

 //fmt.Println(str)
	
	//fmt.Println(pomatch("ab.*c*|", "abab"))
}
