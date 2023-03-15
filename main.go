package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello %s! welcome to the Monkey programming language.\n", user.Username)
	fmt.Printf("feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)

}
