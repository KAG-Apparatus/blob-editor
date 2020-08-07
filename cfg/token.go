package cfg

import "fmt"

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
	NEWLINE   = "\n"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t1 Token) Equal(t2 Token) bool {
	if t1.Literal != t2.Literal || t1.Type != t2.Type {
		return false
	}

	return true
}

func (t Token) String() string {
	return fmt.Sprintf("{%s %s}", t.Type, t.Literal)
}

func LookupWord(word string) TokenType {
	// No keywords currently
	keywords := map[string]TokenType{}
	if keyword, ok := keywords[word]; ok {
		return keyword
	}
	return TEXT
}
