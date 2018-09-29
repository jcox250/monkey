package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	ASTERIX = "*"
	SLASH   = "/"

	// Comparrison
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	// Special Characters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, char byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

// LookupIdent takes an identifier and checks if it is a keyword.
// If it is a keyword then the keywords TokenType is returned.
// If it isn't a keyword then the IDENT token type is returned
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
