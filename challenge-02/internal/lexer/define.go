package lexer

// Lexer converts input characters into predefined language tokens.
// It stores data such as the actual input, current character, current position
// of lexer etc.
type Lexer struct {
	// input stores the input to be processed.
	input []byte
	// char stores the current character that needs to be processed.
	char byte
	// position stores the position of character we are at.
	position int
	// next stores the postion of character next to the current character.
	next int
	// line stores the line number of current character.
	line int
}
