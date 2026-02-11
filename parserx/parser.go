package parserx

import (
	"compiler/tokenx"
	"fmt"
	"os"
)

type Parser struct {
	tokens       []tokenx.Token
	currentIndex int
	currentToken tokenx.Token
	errors       []string
}

func New(tokens []tokenx.Token) *Parser {
	p := &Parser{tokens: tokens}
	if len(tokens) > 0 {
		p.currentToken = tokens[0]
	}
	return p
}

func (p *Parser) nextToken() {
	p.currentIndex++
	if p.currentIndex < len(p.tokens) {
		p.currentToken = p.tokens[p.currentIndex]
	} else {
		p.currentToken = tokenx.Token{Type: tokenx.EOF}
	}
}

func (p *Parser) eat(expected tokenx.TokenType) {
	if p.currentToken.Type == expected {
		p.nextToken()
	} else {
		msg := fmt.Sprintf("Error at position %d: expected %s but got %s (%q)",
			p.currentToken.Pos, expected, p.currentToken.Type, p.currentToken.Lit)
		p.errors = append(p.errors, msg)
		p.nextToken()
	}
}

func (p *Parser) hasErrors() bool {
	return len(p.errors) > 0
}

func (p *Parser) showErrors() {
	fmt.Fprintln(os.Stderr, "=== PARSER ERRORS ===")
	for _, err := range p.errors {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintln(os.Stderr, "=====================")
}

func (p *Parser) parseVarDecl() {
	p.eat(tokenx.KEYWORD_VAR)
	p.eat(tokenx.ID)
	p.eat(tokenx.ASSIGN)
	p.parseExpr()

	if p.currentToken.Type != tokenx.SEMICOLON {
		p.errors = append(p.errors, fmt.Sprintf("Missing ; after var declaration at pos %d", p.currentToken.Pos))
		p.skipTo(tokenx.SEMICOLON, tokenx.RBRACE)
	} else {
		p.eat(tokenx.SEMICOLON)
	}
}

func (p *Parser) parseAssignment() {
	if p.currentToken.Type != tokenx.ID {
		p.errors = append(p.errors, fmt.Sprintf("Expected identifier for assignment at pos %d", p.currentToken.Pos))
		p.nextToken()
		return
	}
	p.eat(tokenx.ID)

	if p.currentToken.Type != tokenx.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("Expected = after identifier at pos %d, got %s", p.currentToken.Pos, p.currentToken.Type))
		p.skipTo(tokenx.SEMICOLON, tokenx.RBRACE)
		return
	}
	p.eat(tokenx.ASSIGN)

	p.parseExpr()

	if p.currentToken.Type != tokenx.SEMICOLON {
		p.errors = append(p.errors, fmt.Sprintf("Missing ; after assignment at pos %d, got %s", p.currentToken.Pos, p.currentToken.Type))
		p.skipTo(tokenx.SEMICOLON, tokenx.RBRACE)
	} else {
		p.eat(tokenx.SEMICOLON)
	}
}

func (p *Parser) skipTo(types ...tokenx.TokenType) {
	typeSet := make(map[tokenx.TokenType]bool)
	for _, t := range types {
		typeSet[t] = true
	}
	for p.currentToken.Type != tokenx.EOF && !typeSet[p.currentToken.Type] {
		p.nextToken()
	}
	if p.currentToken.Type == tokenx.SEMICOLON {
		p.nextToken()
	}
}

func (p *Parser) parseIfStmt() {
	p.eat(tokenx.KEYWORD_IF)
	p.eat(tokenx.LPAREN)
	p.parseCondition()
	p.eat(tokenx.RPAREN)
	p.parseBlock()
}

func (p *Parser) parseWhileStmt() {
	p.eat(tokenx.KEYWORD_WHILE)
	p.eat(tokenx.LPAREN)
	p.parseCondition()
	p.eat(tokenx.RPAREN)
	p.parseBlock()
}

func (p *Parser) parseBlock() {
	p.eat(tokenx.LBRACE)
	for p.currentToken.Type != tokenx.RBRACE && p.currentToken.Type != tokenx.EOF {
		p.parseStatement()
	}
	p.eat(tokenx.RBRACE)
}

func (p *Parser) parseExpr() {
	p.parseTerm()
	for p.currentToken.Type == tokenx.PLUS || p.currentToken.Type == tokenx.MINUS {
		p.nextToken()
		p.parseTerm()
	}
}

func (p *Parser) parseTerm() {
	p.parseFactor()
	for p.currentToken.Type == tokenx.ASTERISK || p.currentToken.Type == tokenx.SLASH {
		p.nextToken()
		p.parseFactor()
	}
}

func (p *Parser) parseFactor() {
	switch p.currentToken.Type {
	case tokenx.INTEGER:
		p.nextToken()
	case tokenx.ID:
		p.nextToken()
	case tokenx.LPAREN:
		p.nextToken()
		p.parseExpr()
		p.eat(tokenx.RPAREN)
	default:
		p.errors = append(p.errors, fmt.Sprintf("Invalid factor at position %d: %s (%q)",
			p.currentToken.Pos, p.currentToken.Type, p.currentToken.Lit))
		p.nextToken()
	}
}

func (p *Parser) parseCondition() {
	p.parseExpr()

	switch p.currentToken.Type {
	case tokenx.LT, tokenx.GT, tokenx.LTE, tokenx.GTE, tokenx.EQ, tokenx.NOT_EQ:
		p.nextToken()
		p.parseExpr()
	default:
		p.errors = append(p.errors, fmt.Sprintf("Expected comparison operator at position %d, got %s",
			p.currentToken.Pos, p.currentToken.Type))
	}
}

func (p *Parser) parseStatement() {
	switch p.currentToken.Type {
	case tokenx.KEYWORD_VAR:
		p.parseVarDecl()
	case tokenx.KEYWORD_IF:
		p.parseIfStmt()
	case tokenx.KEYWORD_WHILE:
		p.parseWhileStmt()
	case tokenx.ID:
		p.parseAssignment()
	case tokenx.LBRACE:
		p.parseBlock()
	default:
		p.errors = append(p.errors, fmt.Sprintf("Invalid statement at %d: %s (%q)",
			p.currentToken.Pos, p.currentToken.Type, p.currentToken.Lit))
		p.nextToken()
	}
}

func (p *Parser) parseProgram() {
	for p.currentToken.Type != tokenx.EOF {
		p.parseStatement()
	}
}

func (p *Parser) Parse() {
	p.parseProgram()

	if p.hasErrors() {
		p.showErrors()
		fmt.Println("Parsing finished with errors")
	} else {
		fmt.Println("Parsing successful - code is syntactically valid")
	}
}
