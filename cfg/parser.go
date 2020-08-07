package cfg

import (
	"fmt"
	"strconv"
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
		Comments:    make([]Comment, 0),
		Assignments: make([]Assignment, 0),
	}
	for !p.curTokenIs(EOF) {
		switch p.curToken.Type {
		case AT:
			fallthrough
		case DOLLAR:
			fallthrough
		case TEXT:
			assignment, err := p.parseAssignment()
			if err != nil {
				return nil, err
			}
			program.Assignments = append(program.Assignments, *assignment)
		default:
			comment, err := p.parseComment()
			if err != nil {
				return nil, err
			}
			program.Comments = append(program.Comments, *comment)
		}
		p.nextToken()
	}
	return program, nil
}

func (p *Parser) parseAssignment() (*Assignment, error) {
	assignment := &Assignment{
		Line: p.lexer.line,
	}

	//Type
	isArray := false
	if p.curTokenIs(AT) && (p.peekTokenIs(TEXT) || p.peekTokenIs(DOLLAR)) {
		isArray = true
		p.nextToken()
		assignment.Type = parseArrayType(p.curToken.Literal)
	} else if p.curTokenIs(TEXT) || p.curTokenIs(DOLLAR) {
		assignment.Type = parseSingleType(p.curToken.Literal)
	} else {
		return nil, fmt.Errorf("invalid token at line %d pos %d: expecting TEXT or DOLLAR got %s", p.lexer.line, p.lexer.position, p.curToken.Type)
	}

	//Identifier
	if !p.peekTokenIs(ASSIGN) {
		p.nextToken()
	}
	assignment.Identifier = p.curToken.Literal
	p.nextToken()

	//Assign
	if !p.curTokenIs(ASSIGN) {
		return nil, fmt.Errorf("invalid token at line %d pos %d: expecting ASSIGN got %s", p.lexer.line, p.lexer.position, p.curToken.Type)
	}
	p.nextToken()

	//Value

	//Multiple values
	if isArray {
		switch assignment.Type {
		case BOOL_ARRAY_TYPE:
			values := make([]bool, 0)
			for p.curToken.Type == TEXT && p.peekToken.Type == SEMICOLON {
				if p.curToken.Literal == "yes" {
					values = append(values, true)
				} else if p.curToken.Literal == "no" {
					values = append(values, false)
				} else {
					return nil, fmt.Errorf("invalid boolean at line %d pos %d: expecting \"yes\" or \"no\" got %s", p.lexer.line, p.lexer.position, p.curToken.Literal)
				}
				p.nextToken()
				p.nextToken()
				p.skipNewLine()
			}
			assignment.Value = values

		case FLOAT_ARRAY_TYPE:
			values := make([]float64, 0)
			for p.curToken.Type == TEXT && p.peekToken.Type == SEMICOLON {
				value, err := strconv.ParseFloat(p.curToken.Literal, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid float at line %d pos %d: %s", p.lexer.line, p.lexer.position, p.curToken.Literal)
				}
				values = append(values, value)
				p.nextToken()
				p.nextToken()
				p.skipNewLine()
			}
			assignment.Value = values

		case INT_ARRAY_TYPE:
			values := make([]int, 0)
			for p.curToken.Type == TEXT && p.peekToken.Type == SEMICOLON {
				value, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid float at line %d pos %d: %s", p.lexer.line, p.lexer.position, p.curToken.Literal)
				}
				values = append(values, int(value))
				p.nextToken()
				p.nextToken()
				p.skipNewLine()
			}
			assignment.Value = values

		case STRING_ARRAY_TYPE:
			fallthrough
		default:
			values := make([]string, 0)
			for p.curToken.Type == TEXT && p.peekToken.Type == SEMICOLON {
				values = append(values, p.curToken.Literal)
				p.nextToken()
				p.nextToken()
				p.skipNewLine()
			}
			assignment.Value = values

		}
		return assignment, nil
	}

	//Single value
	switch assignment.Type {
	case BOOL_TYPE:
		if p.curToken.Literal == "yes" {
			assignment.Value = true
		} else if p.curToken.Literal == "no" {
			assignment.Value = false
		} else {
			return nil, fmt.Errorf("invalid boolean at line %d pos %d: %s", p.lexer.line, p.lexer.position, p.curToken.Literal)
		}
	case FLOAT_TYPE:
		value, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float at line %d pos %d: %s", p.lexer.line, p.lexer.position, p.curToken.Literal)
		}
		assignment.Value = value
	case INT_TYPE:
		value, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid int at line %d pos %d: %s", p.lexer.line, p.lexer.position, p.curToken.Literal)
		}
		assignment.Value = value
	case STRING_ARRAY_TYPE:
		fallthrough
	default:
		assignment.Value = p.curToken.Literal
	}
	p.nextToken()
	return assignment, nil
}

func (p *Parser) parseComment() (*Comment, error) {
	comment := &Comment{
		Line: p.lexer.line,
	}
	start := p.lexer.position
	for !p.curTokenIs(NEWLINE) && !p.curTokenIs(EOF) {
		p.nextToken()
	}
	end := p.lexer.position
	comment.Value = p.lexer.input[start:end]
	return comment, nil
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) skipNewLine() {
	for p.curTokenIs(NEWLINE) {
		p.nextToken()
	}
}

func parseSingleType(literal string) CFGType {
	switch literal {
	case "u8":
		fallthrough
	case "u16":
		fallthrough
	case "u32":
		fallthrough
	case "u64":
		return INT_TYPE
	case "f32":
		fallthrough
	case "f64":
		return FLOAT_TYPE
	case "bool":
		return BOOL_TYPE
	case "$":
		fallthrough
	default:
		return STRING_TYPE
	}
}

func parseArrayType(literal string) CFGType {
	switch literal {
	case "u8":
		fallthrough
	case "u16":
		fallthrough
	case "u32":
		fallthrough
	case "u64":
		return INT_ARRAY_TYPE
	case "f32":
		fallthrough
	case "f64":
		return FLOAT_ARRAY_TYPE
	case "bool":
		return BOOL_ARRAY_TYPE
	case "$":
		fallthrough
	default:
		return STRING_ARRAY_TYPE
	}
}
