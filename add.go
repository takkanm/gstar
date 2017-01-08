package main

import (
	"fmt"
	"io/ioutil"
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
	configFile := configFileName()

	err := ioutil.WriteFile(configFile, []byte(args[0]), 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "write config file is failed.")
		return 1
	}
	fmt.Println("write ", configFile)

	return 0
}
