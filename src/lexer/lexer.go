package lexer

import (
	"github.com/AbdallahSafa/slime/src/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
func (l *Lexer) NextToken() token.Token {
	var tokes token.Token

	switch l.ch {
	case '=':
		tokes = newToken(token.ASSIGN, l.ch)
	case '+':
		tokes = newToken(token.PLUS, l.ch)
	case ';':
		tokes = newToken(token.SEMICOLON, l.ch)
	case '(':
		tokes = newToken(token.LPAREN, l.ch)
	case ')':
		tokes = newToken(token.RPAREN, l.ch)
	case '{':
		tokes = newToken(token.LCURLY, l.ch)
	case '}':
		tokes = newToken(token.RCURLY, l.ch)
	case ',':
		tokes = newToken(token.COMMA, l.ch)
	case 0:
		tokes.Literal = ""
		tokes.Type = token.EOF
	}
	l.readChar()
	return tokes
}
