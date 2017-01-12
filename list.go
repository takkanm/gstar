package main

import (
	"flag"
	"fmt"
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
	allStars := listFlag.Bool("all-stars", false, "show all stars")
	listFlag.Parse(args)

	token := readToken()
	lastPage := *page

	if *allStars {
		lastPage = getStarPageCount(token)
	}
	starLists := getStars(token, *page, lastPage)
	showStars(starLists, func(_ Star) bool {
		return true
	})

	return 0
}

func showStars(starLists []StarList, f func(Star) bool) {
	maxTitleLen := 0
	for _, starList := range starLists {
		stars := starList.stars
		for _, star := range stars {
			titleLen := len(star.FullName)
			if maxTitleLen < titleLen {
				maxTitleLen = titleLen
			}
		}
	}

	for _, starList := range starLists {
		stars := starList.stars

		for _, star := range stars {
			descLen := len(star.Description)
			descSuffix := ""
			if descLen > descLenMax {
				descLen = descLenMax - 3
				descSuffix = "..."
			}
			if f(star) {
				fmt.Printf("%*s : %v\n", maxTitleLen*-1, star.FullName, star.Description[0:descLen]+descSuffix)
			}
		}
	}
}
