package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

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
