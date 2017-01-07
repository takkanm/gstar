package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/mitchellh/cli"
)

func ConfigFileName() string {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(u.HomeDir, ".config", "gstar.config")
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
