package interpreter

import "errors"

//
// ErrEndOfProgram is returned when the end of a program has been reached.
//
var ErrEndOfProgram = errors.New(`End of program`)
