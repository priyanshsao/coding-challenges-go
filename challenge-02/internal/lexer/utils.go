package lexer

func isAlphabet(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func isLiteral(char byte) bool {
	return char == 't' || char == 'f' || char == 'n'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
