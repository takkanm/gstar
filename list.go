package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	token, _ := ioutil.ReadFile(ConfigFileName())
	starList := GetStarList(string(token), 1)
	fmt.Println(starList)
	return 0
}

func GetStarList(token string, page int) string {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.github.com/user/starred?page"+strconv.Itoa(page), nil)
	req.Header.Add("Authorization", "token "+token)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
