package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

type ListStars struct {
}

func (c *ListStars) Synopsis() string {
	return "list stared repository"
}

func (c *ListStars) Help() string {
	return "Usage: gstar list"
}

var descLenMax = 100

func (c *ListStars) Run(args []string) int {
	listFlag := flag.NewFlagSet("list", flag.ExitOnError)
	page := listFlag.Int("page", 1, "show page number")
	listFlag.Parse(args)

	token, _ := ioutil.ReadFile(ConfigFileName())

	ch := make(chan StarList, 1)
	go RequestStars(string(token), *page, ch)

	starList := <-ch
	stars := starList.stars

	maxTitleLen := 0
	for _, star := range stars {
		titleLen := len(star.FullName)
		if maxTitleLen < titleLen {
			maxTitleLen = titleLen
		}
	}

	for _, star := range stars {
		descLen := len(star.Description)
		descSuffix := ""
		if descLen > descLenMax {
			descLen = descLenMax - 3
			descSuffix = "..."
		}

		fmt.Printf("%*s : %v\n", maxTitleLen, star.FullName, star.Description[0:descLen]+descSuffix)
	}

	return 0
}
