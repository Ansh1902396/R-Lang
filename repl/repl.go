package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Ansh1902396/go-interpreter/evaluator"
	"github.com/Ansh1902396/go-interpreter/parser"

	"github.com/Ansh1902396/go-interpreter/lexer"
)

const PROMPT = ">> "

const MONKEY_FACE = `
.--. .-" "-. .--.
/ .. \/ .-. .-. \/ .. \
| | '| / Y \ |' | |
| \ \ \ 0 | 0 / / / |
\ '- ,\.-"""""""-./, -' /
''-' /_ ^ ^ _\ '-''
| \._ _./ |
\ \ '~' / /
'._ '-=-' _.'
'-----'
__,__
`

func Start(in io.Reader, out io.Writer) {
	fmt.Printf(PROMPT)
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
