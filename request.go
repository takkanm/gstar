package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type StarList struct {
	page  int
	stars []Star
}

type Star struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Language    string `json:language`
}

func requestStars(token string, page int, ch chan<- StarList) {
	starList := StarList{page: page}
	starJSON := getStarList(string(token), page)

	json.Unmarshal([]byte(starJSON), &starList.stars)
	ch <- starList
}

func getStarList(token string, page int) string {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.github.com/user/starred?page="+strconv.Itoa(page), nil)
	req.Header.Add("Authorization", "token "+token)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
