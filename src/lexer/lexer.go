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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '?'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumbers() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readIdentifer() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekNext() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tokes token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peekNext() == '=' {
			ch := l.ch
			l.readChar()
			tokes = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tokes = newToken(token.ASSIGN, l.ch)
		}
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
	case '>':
		tokes = newToken(token.GT, l.ch)
	case '<':
		tokes = newToken(token.LT, l.ch)
	case '!':
		if l.peekNext() == '=' {
			ch := l.ch
			l.readChar()
			tokes = token.Token{Type: token.NOTEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tokes = newToken(token.BANG, l.ch)
		}
	case '/':
		tokes = newToken(token.BACKSLASH, l.ch)
	case '*':
		tokes = newToken(token.ASTERISK, l.ch)
	case '-':
		tokes = newToken(token.MINUS, l.ch)
	case 0:
		tokes.Literal = ""
		tokes.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tokes.Literal = l.readIdentifer()
			tokes.Type = token.LookUpIdentifier(tokes.Literal)
			return tokes
		} else if isDigit(l.ch) {
			tokes.Type = token.INTEGER
			tokes.Literal = l.readNumbers()
			return tokes
		} else {
			tokes = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()

	return tokes
}
