package tokenx

type TokenType string

const (
	ILLEGAL    TokenType = "ILLEGAL"
	EOF        TokenType = "EOF"
	ID         TokenType = "ID"
	INTEGER    TokenType = "Integer"
	REAL       TokenType = "Real"
	WHITESPACE TokenType = "WhiteSpace"
)

type Token struct {
	Type TokenType
	Lit  string
	Pos  int
}
