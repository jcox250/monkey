package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jcox250/monkey/lexer"
	"github.com/jcox250/monkey/token"
)

const prompt = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.NewLexer(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%v\n", tok)
		}

	}
}
