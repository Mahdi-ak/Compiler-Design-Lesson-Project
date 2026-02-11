# Simple Compiler - Compiler Design Project

A minimal compiler written in Go with lexer and parser.

## How to Run

### 1. Interactive mode (type code manually)
```bash
go run main.go
```
Type your code, then press **Ctrl+D** (Linux/macOS) or **Ctrl+Z + Enter** (Windows) to finish.

### 2. From a file (recommended)
```bash
go run main.go < test.txt
```
or
```bash
cat test.txt | go run main.go
```

### Example test.txt
```text
var x = 10;
var y = x + 5 * 2;

if (y > 0) {
    var z = y - 3;
    while (z > 0) {
        z = z + 1;
    }
}

x = x + 1;
```

### Output Examples
**Success:**
```
Parsing successful - code is syntactically valid
```

**With error:**
```
=== PARSER ERRORS ===
Missing ; after assignment at pos 42
=====================
Parsing finished with errors
```

## Features
- Lexer: tokenizes input into keywords, identifiers, numbers, operators, punctuation
- Parser: recursive descent syntax checker
- Supports: `var` declarations, assignments, `if`/`while` statements, blocks, arithmetic expressions, comparisons

## Lexer Output Example
When running `go run main.go < test.txt`, the lexer produces tokens like this:

```
Type: KEYWORD_VAR   | Lit: "var"     | Pos: 0
Type: ID            | Lit: "x"       | Pos: 4
Type: ASSIGN        | Lit: "="       | Pos: 6
Type: INTEGER       | Lit: "10"      | Pos: 8
Type: SEMICOLON     | Lit: ";"       | Pos: 10
Type: KEYWORD_VAR   | Lit: "var"     | Pos: 12
Type: ID            | Lit: "y"       | Pos: 16
...
Type: KEYWORD_IF    | Lit: "if"      | Pos: 32
Type: LPAREN        | Lit: "("       | Pos: 35
Type: ID            | Lit: "y"       | Pos: 36
Type: GT            | Lit: ">"       | Pos: 38
Type: INTEGER       | Lit: "0"       | Pos: 40
...
```

Full list of token types produced by the lexer:

| Token Type       | Example              | Description                     |
|------------------|----------------------|---------------------------------|
| KEYWORD_VAR      | `var`                | Variable declaration keyword    |
| KEYWORD_IF       | `if`                 | If statement keyword            |
| KEYWORD_WHILE    | `while`              | While loop keyword              |
| ID               | `x`, `counter`       | Identifier (variable name)      |
| INTEGER          | `42`, `0`            | Integer literal                 |
| ASSIGN           | `=`                  | Assignment operator             |
| PLUS             | `+`                  | Addition                        |
| MINUS            | `-`                  | Subtraction                     |
| ASTERISK         | `*`                  | Multiplication                  |
| SLASH            | `/`                  | Division                        |
| GT, LT, ...      | `>`, `<`, `>=`, ...  | Comparison operators            |
| LPAREN, RPAREN   | `(`, `)`             | Parentheses                     |
| LBRACE, RBRACE   | `{`, `}`             | Braces (blocks)                 |
| SEMICOLON        | `;`                  | Statement terminator            |
| WHITESPACE       | spaces, newlines     | Ignored in parsing              |
| ILLEGAL          | `@`, `#`             | Invalid character               |
| EOF              | —                    | End of input                    |

## Grammar (Simplified EBNF)
```
program         ::= { statement }*
statement       ::= varDecl | assignment | ifStmt | whileStmt | block
varDecl         ::= "var" ID "=" expr ";"
assignment      ::= ID "=" expr ";"
ifStmt          ::= "if" "(" condition ")" block
whileStmt       ::= "while" "(" condition ")" block
block           ::= "{" { statement }* "}"
condition       ::= expr (">" | "<" | ">=" | "<=" | "==" | "!=") expr
expr            ::= term   ( ("+" | "-") term )*
term            ::= factor ( ("*" | "/") factor )*
factor          ::= INTEGER | ID | "(" expr ")"
```

