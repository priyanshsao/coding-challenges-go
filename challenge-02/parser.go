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

	token := p.currentToken
	tokenType := p.currentToken.Type

	switch tokenType {
	case String:
		parsedValue = token.Literal
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

	p.Move()

	if p.currentToken.Type == EndObject {
		return object, nil
	}

	for {
		if p.currentToken.Type != String {
			return nil, fmt.Errorf("unknown token %s at line %d expected key", p.currentToken.Literal, p.lexer.line)
		}

		key := p.currentToken.Literal

		p.Move()

		if p.currentToken.Type != NameSeparator {
			return nil, fmt.Errorf("unknown token %s at line %d expected :", p.currentToken.Literal, p.lexer.line)
		}

		p.Move()

		value, err := p.ParseToken()
		if err != nil {
			return nil, err
		}

		object[key] = value

		p.Move()

		if p.currentToken.Type != ValueSeparator {
			break
		}

		p.Move()
	}

	if p.currentToken.Type != EndObject {
		return nil, fmt.Errorf("unexpected token %s after parsing object at line %d expected }", p.currentToken.Literal, p.lexer.line)
	}

	p.Move()

	return object, nil
}
