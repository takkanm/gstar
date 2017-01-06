package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

type ListStars struct {
}

func (c *ListStars) Synopsis() string {
	return "list stared repository"
}

func (c *ListStars) Help() string {
	return "Usage: gstar list"
}

func (c *ListStars) Run(args []string) int {
	return 0
}

type AddToken struct {
}

func (c *AddToken) Synopsis() string {
	return "add GitHub token"
}

func (c *AddToken) Help() string {
	return "Usage: gstar init github_token"
}

func (c *AddToken) Run(args []string) int {
	return 0
}

func main() {
	c := cli.NewCLI("gstar", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &AddToken{}, nil
		},
		"list": func() (cli.Command, error) {
			return &ListStars{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
