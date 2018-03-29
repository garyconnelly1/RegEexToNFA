
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
