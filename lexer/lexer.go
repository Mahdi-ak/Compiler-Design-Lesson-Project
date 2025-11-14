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

func (l *Lexer) readWhile(pred func(rune) bool) {
	for {
		if l.eof() {
			break
		}
		if !pred(l.peek()) {
			break
		}
		l.next()
	}
}

func (l *Lexer) NextToken() tokenx.Token {
	if l.eof() {
		return tokenx.Token{Type: tokenx.EOF, Lit: "", Pos: l.position}
	}
	ch := l.peek()
	if unicode.IsSpace(ch) {
		l.start = l.position
		l.readWhile(func(r rune) bool { return unicode.IsSpace(r) })
		return l.emit(tokenx.WHITESPACE)
	}
	if unicode.IsLetter(ch) || ch == '_' {
		l.start = l.position
		l.next()
		l.readWhile(func(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' })
		return l.emit(tokenx.ID)
	}
	if unicode.IsDigit(ch) {
		l.start = l.position
		l.readWhile(func(r rune) bool { return unicode.IsDigit(r) })
		if l.peek() == '.' {
			if (l.position+1) <= len(l.input)-1 && unicode.IsDigit(l.input[l.position+1]) {
				l.next()
				l.readWhile(func(r rune) bool { return unicode.IsDigit(r) })
				return l.emit(tokenx.REAL)
			}
			return l.emit(tokenx.INTEGER)
		}
		return l.emit(tokenx.INTEGER)
	}
	if ch == '.' {
		l.start = l.position
		l.next()
		return l.emit(tokenx.ILLEGAL)
	}
	l.start = l.position
	l.next()
	return l.emit(tokenx.ILLEGAL)
}
