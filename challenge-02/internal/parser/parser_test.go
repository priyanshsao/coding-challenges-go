package parser

import (
	"errors"
	"reflect"
	"testing"

	"github.com/priyanshsao/coding-challenges-go/challenge-02/internal/define"
)

func Test_NewParser(t *testing.T) {

	type test struct {
		name     string
		input    string
		expected parser
	}

	tests := []test{
		{
			name:  "Returns parser with current and next token",
			input: "{}",
			expected: parser{
				currentToken: define.Token{define.BeginObject, "{"},
				nextToken:    define.Token{define.EndObject, "}"},
			},
		},
		{
			name:  "Returns Parser with EOF tokens for empty input",
			input: "",
			expected: parser{
				currentToken: define.Token{define.EOF, ""},
				nextToken:    define.Token{define.EOF, ""},
			},
		},
		{
			name:  "Returns Parser with EOF tokens for input with empty lines",
			input: "\n\r\n\t ",
			expected: parser{
				currentToken: define.Token{define.EOF, ""},
				nextToken:    define.Token{define.EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New([]byte(tt.input))
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
		expected parser
	}

	tests := []test{
		{
			name:  "Moves the current Token to next Token",
			input: "{true}",
			expected: parser{
				currentToken: define.Token{define.True, "true"},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := New([]byte(tt.input))
			p.move()

			if p.currentToken != tt.expected.currentToken {
				t.Errorf("expected '%s'(%s) but got '%s'(%s)", tt.expected.currentToken.Literal, tt.expected.currentToken.Type.String(), p.currentToken.Literal, p.currentToken.Type.String())
			}
		})
	}
}

func Test_ParseObject(t *testing.T) {

	type test struct {
		name        string
		input       string
		expected    map[string]any
		expectedErr error
	}

	tests := []test{
		{
			name:  "Should return object for simple input",
			input: `{"key":true}`,
			expected: map[string]any{
				"key": true,
			},
		},
		{
			name:        "Should return error for input without key",
			input:       "{:true}",
			expectedErr: ErrUnknownToken,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := New([]byte(tt.input))
			actualObject, err := p.parseToken()
			if err != nil {
				if tt.expectedErr != nil {
					if !errors.Is(err, tt.expectedErr) {
						t.Fatalf("expected %v but got %v", tt.expectedErr, err)
					}
					return
				}
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actualObject, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, actualObject)
			}

		})
	}
}

func Test_ParseArray(t *testing.T) {

	type test struct {
		name        string
		input       string
		expected    []any
		expectedErr error
	}

	tests := []test{
		{
			name:     "Should return array for simple input",
			input:    `[1,2,"Hi"]`,
			expected: []any{1.0, 2.0, "Hi"},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := New([]byte(tt.input))
			actualArr, err := p.parseToken()
			if err != nil {
				if tt.expectedErr != nil && !errors.Is(err, tt.expectedErr) {
					t.Fatalf("expected %v but got %v", tt.expectedErr, err)
				}
				return
			}

			if !reflect.DeepEqual(actualArr, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, actualArr)
			}
		})
	}
}

func Test_ParseToken(t *testing.T) {

	type test struct {
		name            string
		input           string
		expected        map[string]any
		expectedErrType error
	}

	tests := []test{

		{
			name:  "Should return the parsed JSON map for recursive values",
			input: `{"key1":{"subKey1":{"subKey2":[1,23,4]}}, "key2":[44,6,8]}`,
			expected: map[string]any{
				"key1": map[string]any{
					"subKey1": map[string]any{
						"subKey2": []any{1.0, 23.0, 4.0},
					},
				},
				"key2": []any{44.0, 6.0, 8.0},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := New([]byte(tt.input))

			for {
				parsedObject, err := p.parseToken()
				if err != nil {
					if errors.Is(err, ErrUnexpectedEOF) {
						break
					}

					if tt.expectedErrType != nil && !errors.Is(err, tt.expectedErrType) {
						t.Fatalf("expected %v but got %v", tt.expectedErrType, err)
					}

					return
				}

				if !reflect.DeepEqual(parsedObject, tt.expected) {
					t.Errorf("expected %v but got %v", tt.expected, parsedObject)
				}
			}
		})
	}
}

func Test_Parse(t *testing.T) {

	type test struct {
		name            string
		input           string
		expected        map[string]any
		expectedErrType error
	}

	tests := []test{
		{
			name:            "Should generate unknown token err for multiple JSON objects",
			input:           `{}{}`,
			expectedErrType: ErrUnknownToken,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := New([]byte(tt.input))
			parsedObject, err := p.Parse()
			if err != nil {
				if tt.expectedErrType != nil && !errors.Is(err, tt.expectedErrType) {
					t.Fatalf("expected %v but got %v", tt.expectedErrType, err)
				}
				return
			}

			if !reflect.DeepEqual(parsedObject, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, parsedObject)
			}
		})
	}
}
