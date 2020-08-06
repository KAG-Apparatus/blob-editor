package cfg

type Lexer struct {
	line         int
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
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

func (l *Lexer) NextToken() (tok Token) {
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '$':
		tok = newToken(DOLLAR, l.ch)
	case '#':
		tok = newToken(HASH, l.ch)
	case '@':
		tok = newToken(AT, l.ch)
	case '\n':
		tok = newToken(NEWLINE, l.ch)
		l.line++
		for l.peekChar() == '\n' {
			l.readChar()
		}
	case 0:
		tok = Token{Type: EOF, Literal: ""}
		return
	default:
		if isText(l.ch) {
			word := l.readWord()
			wordType := LookupWord(word)
			tok = Token{Type: wordType, Literal: word}
			return
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return
}

func newToken(tokentype TokenType, ch byte) Token {
	return Token{Type: tokentype, Literal: string(ch)}
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
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
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
