package parser

import (
    "fmt"
    "testing"

    "github.com/Ansh1902396/go-interpreter/lexer"
)

func TestDebugLetReturn(t *testing.T) {
    inputs := []string{
        "let x = 5;",
        "let y = true;",
        "let foobar = y;",
        "return 5;",
        "return true;",
        "return foobar;",
    }

    for _, in := range inputs {
        l := lexer.New(in)
        p := New(l)
        program := p.ParseProgram()
        fmt.Printf("Input: %q\nErrors: %v\nProgram: %#v\n\n", in, p.Errors(), program)
    }
}
