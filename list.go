package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ListStars struct {
}

type StarList struct {
	page  int
	stars []Star
}

type Star struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Language    string `json:language`
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

func RequestStars(token string, page int, ch chan<- StarList) {
	starList := StarList{page: page}
	starJSON := GetStarList(string(token), page)

	json.Unmarshal([]byte(starJSON), &starList.stars)
	ch <- starList
}

func GetStarList(token string, page int) string {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.github.com/user/starred?page="+strconv.Itoa(page), nil)
	req.Header.Add("Authorization", "token "+token)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
