
package main
//Thompsons algorithm
import (
	"fmt"
	
	"os"
	shunting "./ShuntingPackage"
	"io/ioutil"
	"strings"
	"strconv"
	
	
	
)

//"bufio"

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
			//attempt to add the '?' operator
		case '?':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag.initial, edge2: frag.accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
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

func checkFile(){

	count := 0
	i := 0

	for{

	//	reader := bufio.NewReader(os.Stdin)
	var text string
	fmt.Print("Enter egular expression(Enter \"QUIT\" if you wish to exit the program): ")
	fmt.Scan(&text)
	//text, _ := reader.ReadString('\n')
//	fmt.Println("Before trim suffix" + text)
	//trim last 2 ascii characters from the end
	//text = TrimFix(text)

	//check if text equals QUIT
	if text == "QUIT"{
		fmt.Println("Program Ended.")
		os.Exit(2)
	}


	//convert the regular expression from infix to post fix
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
			// st[count] = strconv.Itoa(i)
			fmt.Println("The word" +  " " + word + " appears " + strconv.Itoa(i) + " words in ")
		 }else{
			// fmt.Println("nahh")
		 }
	 }

	 if (count > 0) {
		
		 fmt.Println("Yes! The expression " + text + " exists in the text file " + strconv.Itoa(count) + " times.")

		

	 } else {
		 fmt.Println("No! The expression " + text + " does not exist in the text document.")
	 }
	

	 //re-initialize variables to go through the loop again
	count = 0
	i = 0


	}//end for loop

}

//match string method
func matchString(){
	for{

		//get input for the match string
		var text string
		fmt.Print("Enter egular expression(Enter \"QUIT\" if you wish to exit the program): ")
		fmt.Scan(&text)

		//exit the program if the user entered QUIT
		if text == "QUIT"{
		fmt.Println("Program Ended.")
		os.Exit(2)
		}//end if QUIT

		//convert the regular expression from infix to post fix
		//fmt.Println("intopost + " + shunting.Intopost(text))

		//convert the users string into post fix notation
		expression := shunting.Intopost(text)

		//get user input for the string with which they want to compare the regular expression
		var userString string
		fmt.Print("Enter the string or word you wish to text the regular expression against: ")
		fmt.Scan(&userString)

		//output the true or false result
		fmt.Println(pomatch(expression, userString))
	}//end loop
	 
}//end match string method


func main(){

	//get initial input from the user
	 var input string
	 fmt.Println("Enter \"1\" to check a regular expression against a text file, or \"2\" to see if a regular expression matches an exact string:")
	 fmt.Scan(&input)
	 
	 //if input = 1
	 if input == "1"{
		 //trigger the check file method
		checkFile()
		//if input = 2
	 }else if input =="2"{
		  //trigger the match string method
		  matchString()
		  //else do an error message
		 }else{
		 fmt.Println("Unknown message recieved! Please re run the program and try again.")
	 }	
}
