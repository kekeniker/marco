package main

import (
	"fmt"
	"os"

	"github.com/kekeniker/marco/cmd"
)

func main() {
	execute()
}

func execute() {
	handleError(cmd.Execute(os.Stdout))
}

func handleError(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
