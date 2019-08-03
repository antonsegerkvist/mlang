package interpreter

import "fmt"

//
// Execute takes a program string and executes it.
//
func Execute(program string) error {

	var err error
	var token *Token

	lexer := Lexer{Program: program}
	for token, err = lexer.Next(); err == nil; token, err = lexer.Next() {
		fmt.Printf("==> Token: (%s, |%s|)\n", TokenTranslatorMap[token.Identifier], token.Value)
	}
	if err != ErrEndOfProgram {
		fmt.Println("==> Error: " + err.Error())
	}

	return nil

}
