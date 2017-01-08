package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
)

func configFileName() string {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(u.HomeDir, ".config", "gstar.config")
}

func writeToken(token string) error {
	configFile := configFileName()

	return ioutil.WriteFile(configFile, []byte(token), 0666)
}

func readToken() string {
	token, _ := ioutil.ReadFile(configFileName())
	return string(token)
}
