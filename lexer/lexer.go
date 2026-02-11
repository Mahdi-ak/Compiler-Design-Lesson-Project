package lexer

import (
	"compiler/tokenx"
	"unicode"
)

type Lexer struct {
	input    []rune
	position int
	start    int
	width    int
}

var keywords = map[string]tokenx.TokenType{
	"var":    tokenx.KEYWORD_VAR,
	"func":   tokenx.KEYWORD_FUNC,
	"if":     tokenx.KEYWORD_IF,
	"else":   tokenx.KEYWORD_ELSE,
	"while":  tokenx.KEYWORD_WHILE,
	"return": tokenx.KEYWORD_RETURN,
	"true":   tokenx.KEYWORD_TRUE,
	"false":  tokenx.KEYWORD_FALSE,
}

func New(input string) *Lexer {
	return &Lexer{input: []rune(input)}
}

func (l *Lexer) peek() rune {
	if l.position >= len(l.input) {
		return 0
	}
	return l.input[l.position]
}

func (l *Lexer) next() rune {
	if l.position >= len(l.input) {
		l.width = 0
		return 0
	}
	ch := l.input[l.position]
	l.position++
	l.width = 1
	return ch
}

func (l *Lexer) eof() bool {
	return l.position >= len(l.input)
}

func (l *Lexer) emit(t tokenx.TokenType) tokenx.Token {
	lit := string(l.input[l.start:l.position])
	tok := tokenx.Token{Type: t, Lit: lit, Pos: l.start}
	l.start = l.position
	return tok
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.peek()) && !l.eof() {
		l.next()
	}
	l.start = l.position
}

func lookupKeyword(word string) tokenx.TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return tokenx.ID
}

func (l *Lexer) NextToken() tokenx.Token {
	l.skipWhitespace()

	if l.eof() {
		return tokenx.Token{Type: tokenx.EOF, Lit: "", Pos: l.position}
	}

	ch := l.peek()
	l.start = l.position

	switch {
	case unicode.IsLetter(ch) || ch == '_':
		for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) || l.peek() == '_' {
			l.next()
		}
		literal := string(l.input[l.start:l.position])
		return l.emit(lookupKeyword(literal))

	case unicode.IsDigit(ch):
		for unicode.IsDigit(l.peek()) {
			l.next()
		}
		if l.peek() == '.' {
			l.next()
			for unicode.IsDigit(l.peek()) {
				l.next()
			}
			return l.emit(tokenx.REAL)
		}
		return l.emit(tokenx.INTEGER)

	case ch == '"':
		l.next()
		for !l.eof() && l.peek() != '"' {
			l.next()
		}
		if !l.eof() {
			l.next()
		}
		return l.emit(tokenx.STRING)

	default:
		l.next()

		switch ch {
		case '=':
			if l.peek() == '=' {
				l.next()
				return l.emit(tokenx.EQ)
			}
			return l.emit(tokenx.ASSIGN)

		case '!':
			if l.peek() == '=' {
				l.next()
				return l.emit(tokenx.NOT_EQ)
			}
			return l.emit(tokenx.ILLEGAL)

		case '<':
			if l.peek() == '=' {
				l.next()
				return l.emit(tokenx.LTE)
			}
			return l.emit(tokenx.LT)

		case '>':
			if l.peek() == '=' {
				l.next()
				return l.emit(tokenx.GTE)
			}
			return l.emit(tokenx.GT)

		case '&':
			if l.peek() == '&' {
				l.next()
				return l.emit(tokenx.AND)
			}
			return l.emit(tokenx.ILLEGAL)

		case '|':
			if l.peek() == '|' {
				l.next()
				return l.emit(tokenx.OR)
			}
			return l.emit(tokenx.ILLEGAL)

		case '+':
			return l.emit(tokenx.PLUS)
		case '-':
			return l.emit(tokenx.MINUS)
		case '*':
			return l.emit(tokenx.ASTERISK)
		case '/':
			return l.emit(tokenx.SLASH)

		case '(':
			return l.emit(tokenx.LPAREN)
		case ')':
			return l.emit(tokenx.RPAREN)
		case '{':
			return l.emit(tokenx.LBRACE)
		case '}':
			return l.emit(tokenx.RBRACE)
		case ',':
			return l.emit(tokenx.COMMA)
		case ';':
			return l.emit(tokenx.SEMICOLON)
		case ':':
			return l.emit(tokenx.COLON)

		default:
			return l.emit(tokenx.ILLEGAL)
		}
	}
}
