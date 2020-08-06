package token

const (
	// Symbols
	HASH   = "#"
	DOLLAR = "$"
	AT     = "@"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	TEXT = "TEXT"

	// Operators
	ASSIGN = "="

	// Delimiters
	SEMICOLON = ";"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupWord(word string) TokenType {
	// No keywords currently
	keywords := map[string]TokenType{}
	if keyword, ok := keywords[word]; ok {
		return keyword
	}
	return TEXT
}
