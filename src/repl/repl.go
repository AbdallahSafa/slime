package repl

import (
	"bufio"
	"fmt"
	"github.com/AbdallahSafa/slime/src/lexer"
	"github.com/AbdallahSafa/slime/src/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scnr := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scnr.Scan()
		if !scanned {
			return
		}
		line := scnr.Text()
		l := lexer.New(line)
		for tokes := l.NextToken(); tokes.Type != token.EOF; tokes = l.NextToken() {
			fmt.Printf("%+v\n", tokes)
		}
	}
}
