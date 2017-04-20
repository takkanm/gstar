package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// AddToken is struct for GitHub token manage command
type AddToken struct {
}

// Synopsis is init command synopsis
func (c *AddToken) Synopsis() string {
	return "add GitHub token"
}

// Help is init command usage
func (c *AddToken) Help() string {
	return "Usage: gstar init"
}

// Run is main method for init command
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
