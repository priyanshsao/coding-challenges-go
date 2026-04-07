package main

type Lexer struct {
	input    []byte
	char     byte
	position int
	next     int
	line     int
}

func NewLexer(input []byte) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.Read()
	return l
}

func (l *Lexer) Read() {
	if l.next >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.next]
	}
	l.position = l.next
	l.next++
}

func (l *Lexer) LookNext() byte {
	if l.next >= len(l.input) {
		return 0
	}

	return l.input[l.next]
}

func (l *Lexer) NextToken() Token {
	var token Token

	l.SkipWhiteSpace()

	switch l.char {
	case '[':
		token = Token{BeginArray, string(l.char)}
	case ']':
		token = Token{EndArray, string(l.char)}
	case '{':
		token = Token{BeginObject, string(l.char)}
	case '}':
		token = Token{EndObject, string(l.char)}
	case ':':
		token = Token{NameSeparator, string(l.char)}
	case ',':
		token = Token{ValueSeparator, string(l.char)}
	case '"':
		token = l.ReadString()
	case 0:
		token = Token{EOF, ""}
	default:
		if isDigit(l.char) || l.char == '-' {
			token = l.ReadNum()
		} else if isLiteral(l.char) {
			token = l.ReadLiteral()
		} else {
			token = Token{Illegal, string(l.char)}
		}
	}

	l.Read()

	return token
}

func (l *Lexer) ReadString() Token {
	start := l.position + 1

	for {
		l.Read()

		if l.char == '"' {
			return Token{String, string(l.input[start:l.position])}
		}
		if l.char == 0 {
			return Token{Illegal, string(l.input[start:l.position])}
		}
	}
}

func (l *Lexer) ReadNum() Token {
	start := l.position

	if l.char == '-' {
		l.Read()
	}

	for isDigit(l.char) {
		l.Read()
	}

	if l.char == '.' {
		l.Read()
		for isDigit(l.char) {
			l.Read()
		}
	}

	return Token{Number, string(l.input[start:l.position])}
}

func (l *Lexer) ReadLiteral() Token {
	start := l.position

	for isAlphabet(l.char) {
		l.Read()
	}

	literal := tokenType(l.input[start:l.position])

	switch literal {
	case True:
		return Token{True, "true"}
	case False:
		return Token{False, "false"}
	case Null:
		return Token{Null, "null"}
	default:
		return Token{Illegal, string(literal)}
	}
}

func isAlphabet(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func isLiteral(char byte) bool {
	return char == 't' || char == 'f' || char == 'n'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) SkipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if l.char == '\n' {
			l.line++
		}
		l.Read()
	}
}

func (l *Lexer) CurrentLine() int {
	return l.line
}
