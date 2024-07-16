package main

import (
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

func Connect(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("url:%s failed connect...", url)
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func main() {
	url := "https://www.baidu.com"
	fmt.Println(Connect(url))
}
