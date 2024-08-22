package main

import (
	"fmt"
	"goQuiz/internal/cli"
	"os"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
