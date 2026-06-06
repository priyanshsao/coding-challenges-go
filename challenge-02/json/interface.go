package json

// Parser represents a parser, which parses the tokens into go map.
type Parser interface {
	// Parse parses the tokens from lexer and
	// returns go map with json key-value pairs,
	// if error returns the error.
	Parse() (any, error)
}
