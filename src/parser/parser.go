package parser

import (
	"github.com/AbdallahSafa/slime/src/ast"
	"github.com/AbdallahSafa/slime/src/lexer"
	"github.com/AbdallahSafa/slime/src/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func new(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
