package lexer

import (
	"github.com/jcox250/monkey/token"
)

type Lexer struct {
	input string

	// current position in input (points to current char)
	position int

	// current reading position in input (after current char)
	readPosition int

	// current char under examination
	char byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.char)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.char)
	case '(':
		tok = token.NewToken(token.LPAREN, l.char)
	case ')':
		tok = token.NewToken(token.RPAREN, l.char)
	case ',':
		tok = token.NewToken(token.COMMA, l.char)
	case '+':
		tok = token.NewToken(token.PLUS, l.char)
	case '{':
		tok = token.NewToken(token.LBRACE, l.char)
	case '}':
		tok = token.NewToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	// If we're at the last character of the input i.e the end of the
	// code then assign 0 to char to signify EOF
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		// Otherwise assign the next character in the input to char
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	if l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}
