package main

import (
	"fmt"
)

type Parser struct {
	lexer        *Lexer
	currentToken Token
	nextToken    Token
}

func NewParser(input []byte) *Parser {
	parser := new(Parser)
	lexer := NewLexer(input)

	parser.lexer = lexer
	parser.currentToken = lexer.NextToken()
	parser.nextToken = lexer.NextToken()

	return parser
}

func (p *Parser) Move() {
	p.currentToken = p.nextToken
	p.nextToken = p.lexer.NextToken()
}

func (p *Parser) Parse() (any, error) {
	parsedObject, err := p.ParseToken()
	if err != nil {
		return nil, err
	}

	p.Move()

	if p.currentToken.Type != EOF {
		return nil, fmt.Errorf("expected } but got %s at line %d", p.currentToken.Literal, p.lexer.line)
	}

	return parsedObject, nil
}

func (p *Parser) ParseToken() (any, error) {
	var parsedValue any
	var err error

	switch p.currentToken.Type {
	case BeginObject:
		parsedValue, err = p.ParseObject()
		if err != nil {
			return nil, err
		}
	case EOF:
		return nil, fmt.Errorf("unexpected end of input at line %d", p.lexer.line)
	default:
		return nil, fmt.Errorf("unknown token %s at line %d", p.currentToken.Literal, p.lexer.line)
	}

	return parsedValue, nil
}

func (p *Parser) ParseObject() (any, error) {
	object := make(map[string]any)

	if p.nextToken.Type != EndObject {
		return nil, fmt.Errorf("unexpected token %s after parsing object at line %d", p.currentToken.Literal, p.lexer.line)
	}

	p.Move()

	return object, nil
}
