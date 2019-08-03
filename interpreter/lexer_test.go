package interpreter

import (
	"testing"
)

//
// okProgramList contains an array of ok test programs.
//
var okProgramList = []string{
	`func multiply(x, y) { temp = 0 for x { for y { temp++ } } return temp } temp1 = 0 temp2 = 0 print multiply(temp1, temp2)`,
	`func addition(x, y) { temp = 0 for x { temp++ } for y { temp++ } return temp } temp1 = 0 temp2 = 0 print addtion(temp1, temp2)`,
}

//
// okTokenList contains a list of expected tokens.
//
var okTokenList = [][]Token{
	[]Token{
		Token{Identifier: TokenFunc, Value: "func"},
		Token{Identifier: TokenName, Value: "multiply"},
		Token{Identifier: TokenLParen, Value: "("},
		Token{Identifier: TokenName, Value: "x"},
		Token{Identifier: TokenComma, Value: ","},
		Token{Identifier: TokenName, Value: "y"},
		Token{Identifier: TokenRParen, Value: ")"},
		Token{Identifier: TokenLCBrack, Value: "{"},
		Token{Identifier: TokenName, Value: "temp"},
		Token{Identifier: TokenEqual, Value: "="},
		Token{Identifier: TokenZero, Value: "0"},
		Token{Identifier: TokenFor, Value: "for"},
		Token{Identifier: TokenName, Value: "x"},
		Token{Identifier: TokenLCBrack, Value: "{"},
		Token{Identifier: TokenFor, Value: "for"},
		Token{Identifier: TokenName, Value: "y"},
		Token{Identifier: TokenLCBrack, Value: "{"},
		Token{Identifier: TokenName, Value: "temp"},
		Token{Identifier: TokenInc, Value: "++"},
		Token{Identifier: TokenRCBrack, Value: "}"},
		Token{Identifier: TokenRCBrack, Value: "}"},
		Token{Identifier: TokenReturn, Value: "return"},
		Token{Identifier: TokenName, Value: "temp"},
		Token{Identifier: TokenRCBrack, Value: "}"},
		Token{Identifier: TokenName, Value: "temp1"},
		Token{Identifier: TokenEqual, Value: "="},
		Token{Identifier: TokenZero, Value: "0"},
		Token{Identifier: TokenName, Value: "temp2"},
		Token{Identifier: TokenEqual, Value: "="},
		Token{Identifier: TokenZero, Value: "0"},
		Token{Identifier: TokenPrint, Value: "print"},
		Token{Identifier: TokenName, Value: "multiply"},
		Token{Identifier: TokenLParen, Value: "("},
		Token{Identifier: TokenName, Value: "temp1"},
		Token{Identifier: TokenComma, Value: ","},
		Token{Identifier: TokenName, Value: "temp2"},
		Token{Identifier: TokenRParen, Value: ")"},
	},
	[]Token{
		Token{Identifier: TokenFunc, Value: "func"},
		Token{Identifier: TokenName, Value: "addtion"},
		Token{Identifier: TokenLParen, Value: "("},
		Token{Identifier: TokenName, Value: "x"},
		Token{Identifier: TokenComma, Value: ","},
		Token{Identifier: TokenName, Value: "y"},
		Token{Identifier: TokenRParen, Value: ")"},
		Token{Identifier: TokenLCBrack, Value: "{"},
		Token{Identifier: TokenName, Value: "temp"},
		Token{Identifier: TokenEqual, Value: "="},
		Token{Identifier: TokenZero, Value: "0"},
		Token{Identifier: TokenFor, Value: "for"},
		Token{Identifier: TokenName, Value: "x"},
		Token{Identifier: TokenLCBrack, Value: "{"},
		Token{Identifier: TokenName, Value: "temp"},
		Token{Identifier: TokenInc, Value: "++"},
		Token{Identifier: TokenRCBrack, Value: "}"},
		Token{Identifier: TokenFor, Value: "for"},
		Token{Identifier: TokenName, Value: "y"},
		Token{Identifier: TokenLCBrack, Value: "{"},
		Token{Identifier: TokenName, Value: "temp"},
		Token{Identifier: TokenInc, Value: "++"},
		Token{Identifier: TokenRCBrack, Value: "}"},
		Token{Identifier: TokenReturn, Value: "return"},
		Token{Identifier: TokenName, Value: "temp"},
		Token{Identifier: TokenRCBrack, Value: "}"},
		Token{Identifier: TokenName, Value: "temp1"},
		Token{Identifier: TokenEqual, Value: "="},
		Token{Identifier: TokenZero, Value: "0"},
		Token{Identifier: TokenName, Value: "temp2"},
		Token{Identifier: TokenEqual, Value: "="},
		Token{Identifier: TokenZero, Value: "0"},
		Token{Identifier: TokenPrint, Value: "print"},
		Token{Identifier: TokenName, Value: "addtion"},
		Token{Identifier: TokenLParen, Value: "("},
		Token{Identifier: TokenName, Value: "temp1"},
		Token{Identifier: TokenComma, Value: ","},
		Token{Identifier: TokenName, Value: "temp2"},
		Token{Identifier: TokenRParen, Value: ")"},
	},
}

//
// TestLexer runs a test on the lexer and checks it against the expected output.
//
func TestLexer(t *testing.T) {
	var err error
	var token *Token

	for i, program := range okProgramList {
		lexer := Lexer{Program: program}
		tokens := []*Token{}
		for token, err = lexer.Next(); err == nil; token, err = lexer.Next() {
			tokens = append(tokens, token)
		}
		if err != ErrEndOfProgram {
			t.Errorf("==> Lexer error occured: %s", err.Error())
			return
		}
		if len(okTokenList[i]) != len(tokens) {
			t.Errorf("==> Token count mismatch: %d != %d", len(okTokenList[i]), len(tokens))
			return
		}
		for j, v := range tokens {
			w := okTokenList[i][j]
			if v.Identifier != w.Identifier {
				t.Errorf(
					"==> Token identifier mismatch at location %d: %s != %s",
					j,
					TokenTranslatorMap[v.Identifier],
					TokenTranslatorMap[w.Identifier],
				)
				t.Errorf(
					`==> Token values: %s, %s`,
					v.Value,
					okTokenList[i][j].Value,
				)
				return
			}
		}
	}
}
