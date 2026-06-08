package parser

import (
	"fmt"
	"strconv"

	"github.com/priyanshsao/coding-challenges-go/challenge-02/internal/define"
	"github.com/priyanshsao/coding-challenges-go/challenge-02/internal/lexer"
)

// New returns a new parser that implements the json.Parser interface.
func New(data []byte) *parser {

	p := new(parser)

	p.lexer = lexer.New(data)
	p.nextToken = p.lexer.NextToken()
	p.move()

	return p
}

// Parse parses all the tokens, returns parsed values in go map,
// if error returns error.
func (p *parser) Parse() (any, error) {

	parsedObject, err := p.parseToken()
	if err != nil {
		return nil, err
	}

	p.move()
	if p.currentToken.Type != define.EOF {
		return nil, fmt.Errorf("%w '%s' at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), define.EOF)
	}

	return parsedObject, nil
}

func (p *parser) parseToken() (any, error) {
	var parsedValue any
	var err error

	token := p.currentToken
	tokenType := p.currentToken.Type

	switch tokenType {
	case define.True:
		parsedValue, _ = strconv.ParseBool(token.Literal)
	case define.False:
		parsedValue, _ = strconv.ParseBool(token.Literal)
	case define.Null:
		parsedValue = nil
	case define.Number:
		parsedValue, err = strconv.ParseFloat(token.Literal, 64)
		if err != nil {
			return nil, err
		}
	case define.String:
		parsedValue = token.Literal
	case define.BeginObject:
		parsedValue, err = p.parseObject()
		if err != nil {
			return nil, err
		}
	case define.BeginArray:
		parsedValue, err = p.parseArray()
		if err != nil {
			return nil, err
		}
	case define.EOF:
		return nil, fmt.Errorf("%w at line %d", ErrUnexpectedEOF, p.lexer.CurrentLine())
	case define.Illegal:
		return nil, fmt.Errorf("%w '%s' at line %d", ErrIllegalToken, token.Literal, p.lexer.CurrentLine())
	default:
		return nil, fmt.Errorf("%w '%s' at line %d", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine())
	}

	return parsedValue, nil
}

func (p *parser) parseObject() (any, error) {

	object := make(map[string]any)

	p.move()
	if p.currentToken.Type == define.EndObject {
		return object, nil
	}

	for {
		if p.currentToken.Type != define.String {
			return nil, fmt.Errorf("%w '%s' at line %d, expected key", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine())
		}

		key := p.currentToken.Literal

		p.move()
		if p.currentToken.Type != define.NameSeparator {
			return nil, fmt.Errorf("%w %s at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), define.NameSeparator)
		}

		p.move()
		value, err := p.parseToken()
		if err != nil {
			return nil, err
		}

		object[key] = value

		p.move()
		if p.currentToken.Type != define.ValueSeparator {
			break
		}

		p.move()
	}

	if p.currentToken.Type != define.EndObject {
		return nil, fmt.Errorf("%w '%s' at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), define.EndObject)
	}

	return object, nil
}

func (p *parser) parseArray() (any, error) {

	var array []any

	p.move()
	if p.currentToken.Type == define.EndArray {
		return array, nil
	}

	for {
		parsedValue, err := p.parseToken()
		if err != nil {
			return nil, err
		}

		array = append(array, parsedValue)

		p.move()
		if p.currentToken.Type != define.ValueSeparator {
			break
		}

		p.move()
	}

	if p.currentToken.Type != define.EndArray {
		return nil, fmt.Errorf("%w %s at line %d, expected %v", ErrUnknownToken, p.currentToken.Literal, p.lexer.CurrentLine(), define.EndArray)
	}

	return array, nil
}

func (p *parser) move() {

	p.currentToken = p.nextToken

	if p.nextToken.Type != define.EOF {
		p.nextToken = p.lexer.NextToken()
	}
}
