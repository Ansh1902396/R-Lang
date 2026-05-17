package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Ansh1902396/go-interpreter/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user information:", err)
		return
	}
	fmt.Println("Welcome to the Monkey programming language!", user.Username)

	fmt.Println("Feel free to type in commands and see the results. Type 'exit' to quit.")
	repl.Start(os.Stdin, os.Stdout)
}
