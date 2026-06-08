package parser

import (
	"github.com/priyanshsao/coding-challenges-go/challenge-02/internal/define"
	"github.com/priyanshsao/coding-challenges-go/challenge-02/internal/lexer"
)

// parser converts language tokens into go map.
// It stores a lexer, the current token and next token.
type parser struct {
	// lexer stores the lexer.
	lexer *lexer.Lexer

	// currentToken stores the token being parsed.
	currentToken define.Token

	// nextToken stores the token that is going to be parsed next.
	nextToken define.Token
}
