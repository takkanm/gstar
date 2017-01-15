package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type AddToken struct {
}

func (c *AddToken) Synopsis() string {
	return "add GitHub token"
}

func (c *AddToken) Help() string {
	return "Usage: gstar init"
}

func (c *AddToken) Run(args []string) int {
	token := getGitHubToken()
	err := writeToken(token)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nwrite config file is failed.")
		return 1
	}

	fmt.Println("write token to ", configFileName())
	return 0
}

func getGitHubToken() string {
	fmt.Print("GitHub Token: ")
	token, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	return string(token)
}
