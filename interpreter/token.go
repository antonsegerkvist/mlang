package interpreter

//
// Token identifiers.
//
const (
	TokenLParen = iota
	TokenRParen
	TokenLCBrack
	TokenRCBrack
	TokenComma
	TokenEqual
	TokenInc
	TokenZero
	TokenFor
	TokenFunc
	TokenReturn
	TokenPrint
	TokenName
)

//
// TokenTranslatorMap translates token identifiers to clear.
//
var TokenTranslatorMap = map[uint32]string{
	TokenLParen:  "TOKEN_L_PAREN",
	TokenRParen:  "TOKEN_R_PAREN",
	TokenLCBrack: "TOKEN_LC_BRACK",
	TokenRCBrack: "TOKEN_RC_BRACK",
	TokenComma:   "TOKEN_COMMA",
	TokenEqual:   "TOKEN_EQUAL",
	TokenInc:     "TOKEN_INC",
	TokenZero:    "TOKEN_ZERO",
	TokenFor:     "TOKEN_FOR",
	TokenFunc:    "TOKEN_FUNC",
	TokenReturn:  "TOKEN_RETURN",
	TokenPrint:   "TOKEN_PRINT",
	TokenName:    "TOKEN_NAME",
}

//
// Token contains information about a single token.
//
type Token struct {
	Identifier uint32
	Value      string
}
