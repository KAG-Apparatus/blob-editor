package cfg

import (
	"fmt"
)

type Parser struct {
	lexer *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{
		lexer: lexer,
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseCFGFile() (*CFGFile, error) {
	program := &CFGFile{
		Elements: make([]Element, 0),
	}
	for !p.curTokenIs(EOF) {
		element, err := p.parseElement()
		if err != nil {
			return nil, err
		}
		program.Elements = append(program.Elements, element)
		p.nextToken()
	}
	return program, nil
}

func (p *Parser) parseElement() (Element, error) {
	switch p.curToken.Type {
	case AT:
		fallthrough
	case DOLLAR:
		fallthrough
	case TEXT:
		return p.parseAssignment()
	case NEWLINE:
		return p.parseEmptyLine()
	default:
		return p.parseComment()
	}
}

func (p *Parser) parseAssignment() (Element, error) {
	assignment := &Assignment{
		LeftHand:  make([]Token, 0),
		RightHand: make([]Token, 0),
	}
	for p.curToken.Type != ASSIGN {
		if !p.curTokenIs(AT) && !p.curTokenIs(DOLLAR) && !p.curTokenIs(TEXT) {
			return nil, fmt.Errorf("invalid token at line %d pos %d: expecting TEXT got %s", p.lexer.line+1, p.lexer.position, p.curToken.Type)
		}
		assignment.LeftHand = append(assignment.LeftHand, p.curToken)
		p.nextToken()
	}
	p.nextToken()
	if !p.curTokenIs(TEXT) {
		return nil, fmt.Errorf("invalid token at line %d pos %d: expecting TEXT got %s", p.lexer.line+1, p.lexer.position, p.curToken.Type)
	}
	if p.curToken.Type == TEXT && (p.peekToken.Type == NEWLINE || p.peekToken.Type == EOF) {
		assignment.RightHand = append(assignment.RightHand, p.curToken)
		return assignment, nil
	}
	if p.curToken.Type == TEXT && p.peekToken.Type == SEMICOLON {
		assignment.RightHand = append(assignment.RightHand, p.curToken)
		p.nextToken()
	}
	return assignment, nil
}

func (p *Parser) parseComment() (Element, error) {
	return nil, nil
}

func (p *Parser) parseEmptyLine() (Element, error) {
	return nil, nil
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}
