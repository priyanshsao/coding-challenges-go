package main

import (
	"fmt"
	"strconv"
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
	parser.nextToken = lexer.NextToken()
	parser.Move()

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
		return nil, fmt.Errorf("%w '%s' at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), EOF)
	}

	return parsedObject, nil
}

func (p *Parser) ParseToken() (any, error) {
	var parsedValue any
	var err error

	token := p.currentToken
	tokenType := p.currentToken.Type

	switch tokenType {
	case True:
		parsedValue, _ = strconv.ParseBool(token.Literal)
	case False:
		parsedValue, _ = strconv.ParseBool(token.Literal)
	case Null:
		parsedValue = nil
	case Number:
		parsedValue, err = strconv.ParseFloat(token.Literal, 64)
		if err != nil {
			return nil, err
		}
	case String:
		parsedValue = token.Literal
	case BeginObject:
		parsedValue, err = p.ParseObject()
		if err != nil {
			return nil, err
		}
	case BeginArray:
		parsedValue, err = p.ParseArray()
		if err != nil {
			return nil, err
		}
	case EOF:
		return nil, fmt.Errorf("%w at line %d", ErrUnexpectedEOF, p.lexer.CurrentLine())
	case Illegal:
		return nil, fmt.Errorf("%w '%s' at line %d", ErrIllegalToken, token.Literal, p.lexer.CurrentLine())
	default:
		return nil, fmt.Errorf("%w '%s' at line %d", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine())
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
			return nil, fmt.Errorf("%w '%s' at line %d, expected key", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine())
		}

		key := p.currentToken.Literal

		p.Move()

		if p.currentToken.Type != NameSeparator {
			return nil, fmt.Errorf("%w %s at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), NameSeparator)
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
		return nil, fmt.Errorf("%w '%s' at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), EndObject)
	}

	return object, nil
}

func (p *Parser) ParseArray() (any, error) {
	var array []any

	p.Move()

	if p.currentToken.Type == EndArray {
		return array, nil
	}

	for {
		parsedValue, err := p.ParseToken()
		if err != nil {
			return nil, err
		}

		array = append(array, parsedValue)

		p.Move()

		if p.currentToken.Type != ValueSeparator {
			break
		}

		p.Move()
	}

	if p.currentToken.Type != EndArray {
		return nil, fmt.Errorf("%w %s at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), EndArray)
	}

	return array, nil
}
