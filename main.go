package main

import (
	"bufio"
	"compiler/lexer"
	"compiler/parserx"
	"compiler/tokenx"
	"fmt"
	"io"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var input string
	if (info.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Enter text (Ctrl+D to end):")
		scanner := bufio.NewScanner(os.Stdin)
		var all string
		for scanner.Scan() {
			all += scanner.Text() + "\n"
		}
		input = all
	} else {
		bytes, _ := io.ReadAll(os.Stdin)
		input = string(bytes)
	}

	if input == "" {
		input = "var1 123 45.67 _ok 999. .notok abc_123\n\t42\n"
		fmt.Println("Using sample input:\n" + input)
	}
	tokens := []tokenx.Token{}
	lex := lexer.New(input)
	for {
		tok := lex.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == tokenx.EOF {
			break
		}
		fmt.Printf("Type: %-10s | Lit: %q | Pos: %d\n", tok.Type, tok.Lit, tok.Pos)
	}

	p := parserx.New(tokens)
	p.Parse()
}
