package main

import (
	"fmt"
	"os"
)

type AddToken struct {
}

func (c *AddToken) Synopsis() string {
	return "add GitHub token"
}

func (c *AddToken) Help() string {
	return "Usage: gstar init github_token"
}

func (c *AddToken) Run(args []string) int {
	err := writeToken(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "write config file is failed.")
		return 1
	}

	fmt.Println("write token to ", configFileName())
	return 0
}
