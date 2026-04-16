package main

type tokenType string

// Structural tokens
const (
	BeginArray     tokenType = "["
	BeginObject    tokenType = "{"
	EndArray       tokenType = "]"
	EndObject      tokenType = "}"
	NameSeparator  tokenType = ":"
	ValueSeparator tokenType = ","
)

// Value tokens
const (
	String tokenType = "string"
	Number tokenType = "number"
	True   tokenType = "true"
	False  tokenType = "false"
	Null   tokenType = "null"
)

// Special
const (
	EOF     tokenType = "EOF"
	Illegal tokenType = "ILLEGAL"
)

type Token struct {
	Type    tokenType
	Literal string
}

// For tests to get the string value of token's type
func (tt tokenType) String() string {

	switch tt {
	case BeginObject:
		return "BeginObject"
	case EndObject:
		return "EndObject"
	case BeginArray:
		return "BeginArray"
	case EndArray:
		return "EndArray"
	case NameSeparator:
		return "NameSeperator"
	case ValueSeparator:
		return "ValueSeperator"
	case Number:
		return "Number"
	default:
		if tt == True || tt == False || tt == Null {
			return "Literal"
		}

		return "Unknown"
	}
}
