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
	fmt.Printf(` ______ 
|  ____|
| |__   
|  __|  
| |     
|_|
`)
	fmt.Printf("\nFarcical v0.0.0")
	fmt.Printf("\nShell: %s\n", user.Username)
	fmt.Printf("Enter your command\n")
	repl.Start(os.Stdin, os.Stdout)
}
