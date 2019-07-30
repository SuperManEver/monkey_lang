package token

// Type - contains information about token type
type Type string

const (
	// ILLEGAL - illegal token mark
	ILLEGAL = "ILLEGAL"
	// EOF - eof
	EOF = "EOF"
	// IDENT + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT - int
	INT = "INT"
	// ASSIGN =
	ASSIGN = "="
	// PLUS +
	PLUS = "+"
	// COMMA ,
	COMMA = ","
	// SEMICOLON ;
	SEMICOLON = ";"
	// LPAREN (
	LPAREN = "("

	// RPAREN )
	RPAREN = ")"
	// LBRACE  {
	LBRACE = "{"
	// RBRACE }
	RBRACE = "}"
	// FUNCTION function
	FUNCTION = "FUNCTION"
	// LET let
	LET = "LET"
)

// Token - token struct definition
type Token struct {
	Type    Type
	Literal string
}
