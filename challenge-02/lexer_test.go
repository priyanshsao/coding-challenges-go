package main

import (
	"reflect"
	"testing"
)

func Test_NewLexer(t *testing.T) {
	input := []byte(`{"Key":"value"}`)
	l := NewLexer(input)

	if l.char != '{' {
		t.Errorf("expected '%c' but got '%c'", input[0], l.char)
	}
}

func Test_Read(t *testing.T) {

	type test struct {
		name     string
		input    []byte
		position int
		next     int
		line     int
		expected Lexer
	}

	tests := []test{
		{
			name:  "Read function sets char to input[0]",
			input: []byte("abc"),
			line:  1,
			expected: Lexer{
				char: 'a',
				next: 1,
			},
		},
		{
			name:     "Read function sets char to 0(EOF) at file end",
			input:    []byte("abc"),
			position: 2,
			next:     3,
			line:     1,
			expected: Lexer{
				char:     0,
				position: 3,
				next:     4,
			},
		},
		{
			name:  "Read function sets char to 0(EOF) when input is empty",
			input: []byte(""),
			line:  1,
			expected: Lexer{
				char:     0,
				position: 0,
				next:     1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			l := &Lexer{
				input:    tt.input,
				position: tt.position,
				next:     tt.next,
				line:     tt.line,
			}
			l.Read()

			if l.char != tt.expected.char {
				t.Errorf("expected '%v' but got '%v'", tt.expected.char, l.char)
			}
			if l.position != tt.expected.position {
				t.Errorf("expected %d but got %d", tt.expected.position, l.position)
			}
			if l.next != tt.expected.next {
				t.Errorf("expected %d but got %d", tt.expected.next, l.next)
			}
		})
	}
}

func Test_LookNext(t *testing.T) {

	type test struct {
		name     string
		input    string
		expected byte
	}

	tests := []test{
		{
			name:     "Should return 0 for empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Should return next char",
			input:    "{abc}",
			expected: 'a',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer([]byte(tt.input))
			got := l.LookNext()

			if got != tt.expected {
				t.Errorf("expected '%c' but got '%c'", tt.expected, got)
			}
		})
	}
}

func Test_NextToken(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected []Token
	}

	tests := []test{
		{
			name:  "Should return exact tokens for all base case",
			input: `[]{}:,"HelLo"12trueF`,
			expected: []Token{
				{BeginArray, "["},
				{EndArray, "]"},
				{BeginObject, "{"},
				{EndObject, "}"},
				{NameSeparator, ":"},
				{ValueSeparator, ","},
				{String, "HelLo"},
				{Number, "12"},
				{True, "true"},
				{Illegal, "F"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer([]byte(tt.input))

			for i := 0; i < len(tt.expected); i++ {
				token := l.NextToken()

				if !reflect.DeepEqual(token, tt.expected[i]) {
					t.Errorf("expected '%v' but got '%v'", tt.expected[i], token)
				}
			}
		})
	}
}

func Test_ReadString(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected Token
	}

	tests := []test{
		{
			name:     "returns token for simple string",
			input:    `"key"`,
			expected: Token{String, "key"},
		},
		{
			name:     "returns token for spaced string with space in literal",
			input:    `"   key"`,
			expected: Token{String, "   key"},
		},
		{
			name:     "returns Illegal token for unclosed string",
			input:    `"key`,
			expected: Token{Illegal, "key"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer([]byte(tt.input))
			token := l.ReadString()

			if token.Type != tt.expected.Type {
				t.Fatalf("expected `%s` token type got `%s`", tt.expected.Type, token.Type)
			}
			if token.Literal != tt.expected.Literal {
				t.Errorf("expected token literal to be `%s` but got `%s`", tt.expected.Literal, token.Literal)
			}
		})
	}
}

func Test_ReadNum(t *testing.T) {

	type test struct {
		name     string
		input    string
		expected Token
	}

	tests := []test{
		{
			name:  "Should return token for positive int",
			input: "9004",
			expected: Token{
				Number,
				"9004",
			},
		},
		{
			name:  "Should return token for negative int",
			input: "-236",
			expected: Token{
				Number,
				"-236",
			},
		},
		{
			name:  "Should return token for negative decimal number",
			input: "+23",
			expected: Token{
				Illegal,
				"+",
			},
		},
		{
			name:  "Should return illegal token for non-number",
			input: "a",
			expected: Token{
				Illegal,
				"a",
			},
		},
		{
			name:  "Should return token upto right position",
			input: "4-23.6",
			expected: Token{
				Number,
				"4",
			},
		},
		{
			name:  "Should return illegal token for no digit after '.'",
			input: "423.",
			expected: Token{
				Illegal,
				"423.",
			},
		},
		{
			name:  "Should return illegal token for repititive '.'",
			input: "423..8",
			expected: Token{
				Illegal,
				"423.",
			},
		},
		{
			name:  "Should return token upto valid digits",
			input: "4.23.8",
			expected: Token{
				Number,
				"4.23",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer([]byte(tt.input))

			prevPos := l.position
			numToken := l.NextToken()
			posIncreaseFactor := len(tt.expected.Literal)
			expectedPos := prevPos + posIncreaseFactor

			if !reflect.DeepEqual(numToken, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, numToken)
			}
			if l.position != expectedPos {
				t.Errorf("expected position to be at index %v but got %v", expectedPos, l.position)
			}
		})
	}
}

func Test_ReadLiteral(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected Token
	}

	tests := []test{
		{
			name:  "Should return token for simple literal",
			input: "true",
			expected: Token{
				True,
				"true",
			},
		},
		{
			name:  "Should return illegal token for capitalized literal",
			input: "TRUE",
			expected: Token{
				Illegal,
				"T",
			},
		},
		{
			name:  "Should return illegal token for mix cased literal",
			input: "tRUE",
			expected: Token{
				Illegal,
				"t",
			},
		},
		{
			name:  "Should return illegal token for wrong literal",
			input: "tru",
			expected: Token{
				Illegal,
				"tru",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer([]byte(tt.input))

			prevPos := l.position
			token := l.NextToken()
			posIncreaseFactor := len(tt.expected.Literal)
			expectedPos := prevPos + posIncreaseFactor

			if !reflect.DeepEqual(token, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, token)
			}
			if l.position != expectedPos {
				t.Errorf("expected position to be at index %v but got %v", expectedPos, l.position)
			}
		})
	}
}
