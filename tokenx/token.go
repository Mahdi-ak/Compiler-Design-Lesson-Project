package tokenx

type TokenType string

const (
	ILLEGAL    TokenType = "ILLEGAL"
	EOF        TokenType = "EOF"
	ID         TokenType = "ID"
	INTEGER    TokenType = "Integer"
	REAL       TokenType = "Real"
	WHITESPACE TokenType = "WhiteSpace"

	KEYWORD_VAR    TokenType = "VAR"
	KEYWORD_FUNC   TokenType = "FUNC"
	KEYWORD_IF     TokenType = "IF"
	KEYWORD_ELSE   TokenType = "ELSE"
	KEYWORD_WHILE  TokenType = "WHILE"
	KEYWORD_RETURN TokenType = "RETURN"
	KEYWORD_TRUE   TokenType = "TRUE"
	KEYWORD_FALSE  TokenType = "FALSE"

	STRING   TokenType = "STRING"
	ASSIGN   TokenType = "ASSIGN"
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	ASTERISK TokenType = "ASTERISK"
	SLASH    TokenType = "SLASH"
	EQ       TokenType = "EQ"
	NOT_EQ   TokenType = "NOT_EQ"
	LT       TokenType = "LT"
	GT       TokenType = "GT"
	LTE      TokenType = "LTE"
	GTE      TokenType = "GTE"
	AND      TokenType = "AND"
	OR       TokenType = "OR"

	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	COLON     TokenType = "COLON"
)

type Token struct {
	Type TokenType
	Lit  string
	Pos  int
}
