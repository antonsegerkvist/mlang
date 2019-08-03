package interpreter

import (
	"bytes"
	"errors"
	"unicode"
)

//
// Lexer contains the lexer methods.
//
type Lexer struct {

	//
	// Program contains a unhandled part of the program.
	//

	Program string
}

//
// Peek returns the rune at the current location.
//
func (lexer *Lexer) Peek() rune {
	runes := []rune(lexer.Program)
	if len(runes) > 0 {
		return runes[0]
	}
	return '\000'
}

//
// Get returns the rune at the current location and advances lexer.
//
func (lexer *Lexer) Get() (rune, error) {
	runes := []rune(lexer.Program)
	if len(runes) > 0 {
		var ret rune
		ret, lexer.Program = runes[0], string(runes[1:])
		return ret, nil
	}
	return '\000', ErrEndOfProgram
}

//
// Identifier parses the program until a identifier is found.
//
func (lexer *Lexer) Identifier(r rune) (*Token, error) {
	var err error
	buffer := bytes.Buffer{}
	if unicode.IsLetter(r) {
		buffer.WriteRune(r)
		for r = lexer.Peek(); unicode.IsLetter(r) || unicode.IsNumber(r); r = lexer.Peek() {
			r, err = lexer.Get()
			if err != nil {
				return nil, err
			}
			buffer.WriteRune(r)
		}
		switch buffer.String() {

		case "for":
			return &Token{
				Identifier: TokenFor,
				Value:      buffer.String(),
			}, nil

		case "func":
			return &Token{
				Identifier: TokenFunc,
				Value:      buffer.String(),
			}, nil

		case "print":
			return &Token{
				Identifier: TokenPrint,
				Value:      buffer.String(),
			}, nil

		case "return":
			return &Token{
				Identifier: TokenReturn,
				Value:      buffer.String(),
			}, nil

		default:
			return &Token{
				Identifier: TokenName,
				Value:      buffer.String(),
			}, nil

		}
	}
	return nil, errors.New(`Bad identifier: ` + string(r))
}

//
// Next returns the next token from the program.
//
func (lexer *Lexer) Next() (*Token, error) {
	var r rune
	var err error

	for r, err = lexer.Get(); unicode.IsSpace(r) && err == nil; r, err = lexer.Get() {
	}
	if err != nil {
		return nil, err
	}

	switch r {

	case '(':
		return &Token{
			Identifier: TokenLParen,
			Value:      string(r),
		}, nil

	case ')':
		return &Token{
			Identifier: TokenRParen,
			Value:      string(r),
		}, nil

	case '{':
		return &Token{
			Identifier: TokenLCBrack,
			Value:      string(r),
		}, nil

	case '}':
		return &Token{
			Identifier: TokenRCBrack,
			Value:      string(r),
		}, nil

	case '0':
		return &Token{
			Identifier: TokenZero,
			Value:      string(r),
		}, nil

	case '=':
		return &Token{
			Identifier: TokenEqual,
			Value:      string(r),
		}, nil

	case ',':
		return &Token{
			Identifier: TokenComma,
			Value:      string(r),
		}, nil

	case '+':
		if lexer.Peek() == '+' {
			tmp, _ := lexer.Get()
			return &Token{
				Identifier: TokenInc,
				Value:      string(r) + string(tmp),
			}, nil
		}
		return nil, errors.New(`Bad identifier: ` + string(r))

	default:
		return lexer.Identifier(r)

	}
}

// //
// // Lexer returns a list of lexical tokens from the specified string.
// //
// func Lexer(program string) (*[]Token, error) {

// 	tokens := []Token{}
// 	buffer := &bytes.Buffer{}

// 	for _, v := range []rune(program + " ") {

// 		switch buffer.String() {

// 		case "(":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenLParen,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case ")":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenRParen,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "{":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenLCBrack,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "}":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenRCBrack,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case ",":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenComma,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "=":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenEqual,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "+":
// 			if string(v) == "+" {
// 				tokens = append(tokens, Token{
// 					Identifier: TokenInc,
// 					Value:      buffer.String() + string(v),
// 				})
// 				buffer.Reset()
// 			} else {
// 				return nil, errors.New("Unsupported token +" + string(v))
// 			}

// 		case "0":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenZero,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "for":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenFor,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "func":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenFunc,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "return":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenReturn,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		case "print":
// 			tokens = append(tokens, Token{
// 				Identifier: TokenPrint,
// 				Value:      buffer.String(),
// 			})
// 			buffer.Reset()
// 			if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		default:
// 			if buffer.Len() > 0 && isNameTerminator(v) {
// 				tokens = append(tokens, Token{
// 					Identifier: TokenName,
// 					Value:      buffer.String(),
// 				})
// 				buffer.Reset()
// 				if !unicode.IsSpace(v) {
// 					buffer.WriteRune(v)
// 				}
// 			} else if !unicode.IsSpace(v) {
// 				buffer.WriteRune(v)
// 			}

// 		}
// 	}

// 	return &tokens, nil

// }

// //
// // isNameTerminator returns true if the character specifies a name termination
// // character.
// //
// func isNameTerminator(v rune) bool {
// 	return false ||
// 		unicode.IsSpace(v) ||
// 		string(v) == "(" ||
// 		string(v) == ")" ||
// 		string(v) == "{" ||
// 		string(v) == "}" ||
// 		string(v) == "," ||
// 		string(v) == "=" ||
// 		string(v) == "+" ||
// 		string(v) == "0"
// }
