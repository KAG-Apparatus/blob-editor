package lexer

import (
	"github.com/KAG-Apparatus/blob-editor/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if inputLength := len(l.input); l.readPosition >= inputLength {
		l.ch = 0
		l.position = l.readPosition
		return
	}
	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '$':
		tok = newToken(token.DOLLAR, l.ch)
	case '#':
		tok = newToken(token.HASH, l.ch)
	case '@':
		tok = newToken(token.AT, l.ch)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
		return
	default:
		if isText(l.ch) {
			word := l.readWord()
			wordType := token.LookupWord(word)
			tok = token.Token{Type: wordType, Literal: word}
			return
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return
}

func newToken(tokentype token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokentype, Literal: string(ch)}
}

func (l *Lexer) readWord() string {
	position := l.position
	for isText(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isText(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' ||
		('0' <= ch && ch <= '9') ||
		'-' == ch || '.' == ch || '*' == ch
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
