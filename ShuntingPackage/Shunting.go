//shunting yard algorithm

package ShuntingPackage

import (
	
)
//"fmt"
/*
	The main knowledge for this function came from our learn online moodle page. - https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e .
*/
//function for changing infix to post fix notation
func Intopost(infix string)string{
	//initialize variables needed
	specials := map[rune]int{'*':10, '.': 9, '|':8}
	pofix := []rune{}
	s := []rune{} 

//for each character in the expression
	for _, r := range infix{
		switch {
			//if the character = '(', simply append that character onto the array s.
		case r == '(' :
			s = append(s,r)

			//if the character = ')', while the last index in the s array, append the last index of the s array to the pofix array, 
			//and clear the s array up to the value in the last index of that array.
		case r == ')' :
			for(s[len(s)-1] != '('){
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]  
			}//end for
			//equate the array s to just the last index in that array
			s = s[:len(s)-1]  

			//if the character = one of the desegnated characters in the specials map,
		case specials[r] > 0 :
			//for when the lenght of array s in greater than 0 and specials array at index r(the current character),
			//is less than or = to the value at the specials array at index of the value at the last index of array s,
			for(len(s) > 0 && specials[r] <= specials[s[len(s)-1]] ){
				//append the last index of array s to pofix, and clear array s up to the last index
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] 
			}//end for
			//append r(current character to array s.)
			s = append(s,r)
		default:
			//if the current character does not meet one of the above cases, append the current character to pofix array.
		pofix = append(pofix,r)
		}
	
	}//end for range

	for len(s) > 0 {
	pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] 
}


	return string(pofix)
}
