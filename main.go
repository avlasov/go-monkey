package main

import (
	"fmt"
	"github.com/avlasov/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type commnads\n")
	repl.Start(os.Stdin, os.Stdout)
}
