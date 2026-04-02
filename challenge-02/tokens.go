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
