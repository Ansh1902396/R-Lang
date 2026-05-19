package runner

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Ansh1902396/go-interpreter/evaluator"
	"github.com/Ansh1902396/go-interpreter/lexer"
	"github.com/Ansh1902396/go-interpreter/object"
	"github.com/Ansh1902396/go-interpreter/parser"
)

const SourceExtension = ".rlang"

func RunFile(path string, errOut io.Writer) int {
	if filepath.Ext(path) != SourceExtension {
		fmt.Fprintf(errOut, "unsupported file extension %q; expected %q\n", filepath.Ext(path), SourceExtension)
		return 1
	}

	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(errOut, "could not read %s: %v\n", path, err)
		return 1
	}

	evaluated, parserErrors := Run(string(input))
	if len(parserErrors) != 0 {
		fmt.Fprintf(errOut, "parser errors in %s:\n", path)
		for _, msg := range parserErrors {
			fmt.Fprintf(errOut, "\t%s\n", msg)
		}
		return 1
	}

	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		errObj := evaluated.(*object.Error)
		fmt.Fprintf(errOut, "runtime error in %s: %s\n", path, errObj.Message)
		return 1
	}

	return 0
}

func Run(input string) (object.Object, []string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return nil, p.Errors()
	}

	env := object.NewEnvironment()
	return evaluator.Eval(program, env), nil
}
