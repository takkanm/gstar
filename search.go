package main

import "strings"

// SearchStars is struct for search command
type SearchStars struct {
}

// Synopsis is search command synopsis
func (c *SearchStars) Synopsis() string {
	return "search stared repository"
}

// Help is search command help
func (c *SearchStars) Help() string {
	return "Usage: gstar search"
}

// Run is main method for search command
func (c *SearchStars) Run(args []string) int {
	token := readToken()
	lastPage := getStarPageCount(token)

	starLists := getStars(token, 1, lastPage)

	showStars(starLists, func(s Star) bool {
		return strings.Contains(s.FullName, args[0]) || strings.Contains(s.Description, args[0])
	})

	return 0
}
