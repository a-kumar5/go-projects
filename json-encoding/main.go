package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	log.Println(response)
	if response.s
}
