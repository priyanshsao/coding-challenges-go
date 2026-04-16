package main

import (
	"testing"
)

func Test_NewParser(t *testing.T) {

	type test struct {
		name     string
		input    string
		expected Parser
	}

	tests := []test{
		{
			name:  "Returns parser with current and next token",
			input: "{}",
			expected: Parser{
				currentToken: Token{BeginObject, "{"},
				nextToken:    Token{EndObject, "}"},
			},
		},
		{
			name:  "Returns Parser with EOF tokens for empty input",
			input: "",
			expected: Parser{
				currentToken: Token{EOF, ""},
				nextToken:    Token{EOF, ""},
			},
		},
		{
			name:  "Returns Parser with EOF tokens for input with empty lines",
			input: "\n\r\n\t ",
			expected: Parser{
				currentToken: Token{EOF, ""},
				nextToken:    Token{EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser([]byte(tt.input))
			pExpected := tt.expected

			if p.currentToken != pExpected.currentToken {
				t.Errorf("expected '%s'(%s) but got '%s'(%s)", pExpected.currentToken.Literal, pExpected.currentToken.Type.String(), p.currentToken.Literal, p.currentToken.Type.String())
			}
			if p.nextToken != pExpected.nextToken {
				t.Errorf("expected '%s'(%s) but got '%s'(%s)", pExpected.nextToken.Literal, pExpected.nextToken.Type.String(), p.nextToken.Literal, p.nextToken.Type.String())
			}
		})
	}
}

func Test_Move(t *testing.T) {

	type test struct {
		name     string
		input    string
		expected Parser
	}

	tests := []test{
		{
			name:  "Moves the current Token to next Token",
			input: "{true}",
			expected: Parser{
				currentToken: Token{True, "true"},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := NewParser([]byte(tt.input))
			p.Move()

			if p.currentToken != tt.expected.currentToken {
				t.Errorf("expected '%s'(%s) but got '%s'(%s)", tt.expected.currentToken.Literal, tt.expected.currentToken.Type.String(), p.currentToken.Literal, p.currentToken.Type.String())
			}
		})
	}
}
