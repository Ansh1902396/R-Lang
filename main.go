package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Ansh1902396/go-interpreter/repl"
	"github.com/Ansh1902396/go-interpreter/runner"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [file%s]\n", os.Args[0], runner.SourceExtension)
		os.Exit(1)
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Fprintf(os.Stdout, "usage: %s [file%s]\n", os.Args[0], runner.SourceExtension)
			return
		}
		os.Exit(runner.RunFile(os.Args[1], os.Stderr))
	}

	user, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user information:", err)
		return
	}
	fmt.Println("Welcome to the Monkey programming language!", user.Username)

	fmt.Println("Feel free to type in commands and see the results. Type 'exit' to quit.")
	repl.Start(os.Stdin, os.Stdout)
}
