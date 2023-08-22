package main

import (
	"farcical/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -- Farcical Shell\n", user.Username)
	fmt.Printf("Enter your command\n")
	repl.Start(os.Stdin, os.Stdout)
}
