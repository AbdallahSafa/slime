package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INTEGER    = "INTEGER"    // 1343456

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LCURLY = "{"
	RCURLY = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"func": FUNCTION,
	"let":  LET,
}

func LookUpIdentifier(identity string) TokenType {
	if tokes, ok := keywords[identity]; ok {
		return tokes
	}
	return IDENTIFIER
}

type Token struct {
	Type    TokenType
	Literal string
}
