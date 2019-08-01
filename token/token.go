package token

// Type definition
type Type string

const (
	// ILLEGAL const
	ILLEGAL = "ILLEGAL"

	// EOF const
	EOF = "EOF"

	// IDENT const
	IDENT = "IDENT" // add, foobar, x, y, ...

	// INT const
	INT = "INT" // 1343456

	// ASSIGN const
	ASSIGN = "="

	// PLUS const
	PLUS = "+"

	// MINUS const
	MINUS = "-"

	// BANG const
	BANG = "!"

	// ASTERISK const
	ASTERISK = "*"

	// SLASH const
	SLASH = "/"

	// LT lt
	LT = "<"

	// GT gt
	GT = ">"

	// EQ eq
	EQ = "=="

	// NOT_EQ not eq
	NOT_EQ = "!="

	// COMMA ,
	COMMA = ","

	// SEMICOLON ;
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Token type
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checking for identifier
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
