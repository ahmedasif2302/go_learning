package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Posts struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Id     int    `json:"id"`
}

func goApiCallMain() {
	fmt.Println("Making api call in go")
	makeGetApiCall()
}

func makeGetApiCall() {
	var url = "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// always need to close the api call connection at the end of the func executions
	defer response.Body.Close()

	// read all the json content from the body to string as it go does not support json format
	content, _ := io.ReadAll(response.Body)

	var posts []Posts

	// convert json to slices
	json.Unmarshal(content, &posts)
	jsonPosts, _ := json.MarshalIndent(posts, "", "\t")
	fmt.Println(string(jsonPosts))
}

func makePostApiCall() {}
