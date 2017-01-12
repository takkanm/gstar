package main

import "strings"

type SearchStars struct {
}

func (c *SearchStars) Synopsis() string {
	return "search stared repository"
}

func (c *SearchStars) Help() string {
	return "Usage: gstar search"
}

func (c *SearchStars) Run(args []string) int {
	token := readToken()
	lastPage := getStarPageCount(token)

	starLists := getStars(token, 1, lastPage)

	showStars(starLists, func(s Star) bool {
		return strings.Contains(s.FullName, args[0]) || strings.Contains(s.Description, args[0])
	})

	return 0
}
