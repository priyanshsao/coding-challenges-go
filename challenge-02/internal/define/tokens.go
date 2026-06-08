package define

type Token struct {
	// Type represents token type: structural, value, special etc.
	Type TokenType

	// Literal represents the literal value stored for the token.
	Literal string
}

// TokenType denotes tokens, can be one of this :structural, value, special.
type TokenType string

const (
	// BeginArray denotes starting of array.
	BeginArray TokenType = "["

	// BeginObject denotes starting of json-object.
	BeginObject TokenType = "{"

	// EndArray denotes end of an array.
	EndArray TokenType = "]"

	// EndObject denotes end of json-object.
	EndObject TokenType = "}"

	// NameSeperator denotes separation between
	// json key-value pairs.
	NameSeparator TokenType = ":"

	// ValueSeperator denotes separation between
	// elements in json-object.
	ValueSeparator TokenType = ","
)

const (
	// String denotes a normal string value.
	String TokenType = "string"

	// Number denotes numbers : int, float etc.
	Number TokenType = "number"

	// True denotes true literal.
	True TokenType = "true"

	// False denotes false literal.
	False TokenType = "false"

	// Null denotes null value (acc to json-spec).
	Null TokenType = "null"
)

const (
	// EOF denotes the end of file.
	EOF TokenType = "EOF"

	// Illegal denotes unkown charcters,
	// that are not defined in language grammer.
	Illegal TokenType = "ILLEGAL"
)

func (tt TokenType) String() string {

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
	case String:
		return "String"
	case Number:
		return "Number"
	case EOF:
		return "EOF"
	default:
		if tt == True || tt == False || tt == Null {
			return "Literal"
		}

		return "Unknown"
	}
}
