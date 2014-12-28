package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

func FetchPage(page string) (string, error) {
	url := page
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func FindTag(body, open, close string) (string, error) {
	e := open + "(.*?)" + close
	r, err := regexp.Compile(e)
	if err != nil {
		return "", err
	}
	m := r.FindStringSubmatch(body)
	if m == nil {
		return "", errors.New("Not found.")
	}
	return regexp.MustCompile("<.*?>").ReplaceAllString(m[1], ""), nil
}
