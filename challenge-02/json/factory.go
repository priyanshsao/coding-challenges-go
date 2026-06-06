package json

import (
	"github.com/priyanshsao/coding-challenges-go/challenge-02/internal/parser"
)

// NewParser returns new parser interface.
func NewParser(data []byte) Parser {

	return parser.New(data)
}
