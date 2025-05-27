package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d</h1>", <-ch)
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func generate(limit int, ch chan<- int) {
	for i := 2; i <= limit; i++ {
		ch <- i
	}
	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst)
}

func sieve(limit int) {
	ch := make(chan int)
	go generate(limit, ch)
	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		fmt.Print(prime, " ")
	}
}

func main() {
	sieve(100)
	/*
		var nextID nextCh = make(chan int)
		go counter(nextID)
		http.HandleFunc("/", nextID.handler)
		log.Fatal(http.ListenAndServe(":8080", nil))


		   results := make(chan result)

		   	list := []string{
		   		"https://amazon.com",
		   		"https://google.com",
		   		"https://facebook.com",
		   		"https://x.com",
		   	}

		   	for _, url := range list {
		   		go get(url, results)
		   	}

		   	for range list {
		   		r := <-results
		   		if r.err != nil {
		   			log.Printf("%-20s %s\n", r.url, r.err)
		   		} else {
		   			log.Printf("%-20s %s\n", r.url, r.latency)
		   		}
		   	}
	*/
}
