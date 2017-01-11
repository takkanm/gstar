package main

type SearchStars struct {
}

func (c *SearchStars) Synopsis() string {
	return "search stared repository"
}

func (c *SearchStars) Help() string {
	return "Usage: gstar search"
}

func (c *SearchStars) Run(args []string) int {

	return 0
}
