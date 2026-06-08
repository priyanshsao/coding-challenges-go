package lexer

import (
	"github.com/priyanshsao/coding-challenges-go/challenge-02/internal/define"
	"github.com/sirupsen/logrus"
)

// New returns a new lexer.
func New(input []byte) *Lexer {

	l := new(Lexer)
	l.input = input
	l.line = 1
	l.read()

	return l
}

// NextToken returns token, by processing the current character.
func (l *Lexer) NextToken() define.Token {

	var token define.Token

	l.skipWhiteSpace()

	switch l.char {
	case '[':
		token = define.Token{Type: define.BeginArray, Literal: string(l.char)}
	case ']':
		token = define.Token{Type: define.EndArray, Literal: string(l.char)}
	case '{':
		token = define.Token{Type: define.BeginObject, Literal: string(l.char)}
	case '}':
		token = define.Token{Type: define.EndObject, Literal: string(l.char)}
	case ':':
		token = define.Token{Type: define.NameSeparator, Literal: string(l.char)}
	case ',':
		token = define.Token{Type: define.ValueSeparator, Literal: string(l.char)}
	case '"':
		token = l.readString()
	case 0:
		token = define.Token{Type: define.EOF, Literal: ""}
	default:
		if isDigit(l.char) || l.char == '-' {
			token = l.readNum()
		} else if isLiteral(l.char) {
			token = l.readLiteral()
		} else {
			token = define.Token{Type: define.Illegal, Literal: string(l.char)}
		}
	}

	logrus.Debugf("Token generated: %v(%q)", token.Type, token.Literal)

	l.read()

	return token
}

// CurrentLine returns the line number which
// is being processed by the lexer.
func (l *Lexer) CurrentLine() int {

	return l.line
}

func (l *Lexer) read() {

	if l.next >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.next]
	}

	l.position = l.next
	l.next++
}

func (l *Lexer) lookNext() byte {

	if l.next >= len(l.input) {
		return 0
	}

	return l.input[l.next]
}

func (l *Lexer) skipWhiteSpace() {

	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if l.char == '\n' {
			l.line++
		}

		l.read()
	}
}

func (l *Lexer) readString() define.Token {

	start := l.position + 1

	for {
		l.read()

		if l.char == '"' {
			return define.Token{Type: define.String, Literal: string(l.input[start:l.position])}
		}
		if l.char == 0 {
			return define.Token{Type: define.Illegal, Literal: string(l.input[start:l.position])}
		}
	}
}

func (l *Lexer) readNum() define.Token {

	start := l.position

	for isDigit(l.lookNext()) {
		l.read()
	}

	if l.lookNext() == '.' {
		l.read()
		if !isDigit(l.lookNext()) {
			return define.Token{Type: define.Illegal, Literal: string(l.input[start : l.position+1])}
		}
		for isDigit(l.lookNext()) {
			l.read()
		}
	}

	return define.Token{
		Type:    define.Number,
		Literal: string(l.input[start : l.position+1]),
	}
}

func (l *Lexer) readLiteral() define.Token {

	start := l.position

	for isAlphabet(l.lookNext()) {
		l.read()
	}

	literal := define.TokenType(l.input[start : l.position+1])

	switch literal {
	case define.True:
		return define.Token{Type: define.True, Literal: "true"}
	case define.False:
		return define.Token{Type: define.False, Literal: "false"}
	case define.Null:
		return define.Token{Type: define.Null, Literal: "nil"}
	default:
		return define.Token{Type: define.Illegal, Literal: string(literal)}
	}
}
