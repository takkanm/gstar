package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// StarList is struct for GitHub Star page
type StarList struct {
	page  int
	stars []Star
}

// Star is struct for GitHub Star
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

func getStars(token string, firstPage int, lastPage int) []StarList {
	requestPageSize := lastPage - firstPage + 1
	results := make([]StarList, requestPageSize)
	ch := make(chan StarList, 1)

	for i := firstPage; i <= lastPage; i++ {
		go requestStars(token, i, ch)
	}

	for i := 0; i < requestPageSize; i++ {
		starList := <-ch
		results = append(results, starList)
	}

	return results
}

func getStarPageCount(token string) int {
	client := &http.Client{}

	req, _ := http.NewRequest("HEAD", "https://api.github.com/user/starred", nil)
	req.Header.Add("Authorization", "token "+token)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	linkHeader := resp.Header["Link"][0]
	lastURLPart := strings.Split(linkHeader, ",")[1]
	lastURLString := strings.Trim(strings.Split(lastURLPart, ";")[0], " <>")
	u, _ := url.Parse(lastURLString)
	q := u.Query()
	page, _ := strconv.Atoi(q["page"][0])

	return page
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
