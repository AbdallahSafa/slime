package main

import (
	"fmt"
	"github.com/AbdallahSafa/slime/src/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Username)
	fmt.Println(user.HomeDir)
	repl.Start(os.Stdin, os.Stdout)
}
