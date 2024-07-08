package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INTEGER    = "INTEGER"    // 1343456

	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	BACKSLASH = "/"
	LT        = "<"
	GT        = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LCURLY = "{"
	RCURLY = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
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
