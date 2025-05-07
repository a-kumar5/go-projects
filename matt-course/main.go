package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(len(p))
	return l, nil
}

func main() {
	// class 18 methods and interfaces
	var b ByteCounter
	f1, _ := os.Open("test.txt")
	//f2, _ := os.Create("out.txt")
	f2 := &b

	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
	fmt.Println(b)

	//fmt.Println(hello.Say(os.Args[1:]))
	// class 3 code
	/*
		var sum float64
		var n int

		for {
			var val float64
			_, err := fmt.Fscanln(os.Stdin, &val)
			if err != nil {
				break
			}

			sum += val
			n++
		}

		if n == 0 {
			fmt.Fprintln(os.Stderr, "no values")
			os.Exit(-1)
		}

		fmt.Println("The average is", sum/float64(n))
	*/

	//class 4:= Strings
	// program to search and replace word in a file
	/*
			if len(os.Args) < 3 {
				fmt.Fprintf(os.Stderr, "No enough input args")
				os.Exit(-1)
			}

			old, new := os.Args[1], os.Args[2]
			fmt.Println(old)
			fmt.Println(new)

			scan := bufio.NewScanner(os.Stdin)
			//var line string

			for scan.Scan() {
				//fmt.Println(scan.Text())
				s := strings.Split(scan.Text(), old)
				//fmt.Println(len(s))
				r := strings.Join(s, new)
				fmt.Println(r)
			}

		// class 05
		scan := bufio.NewScanner(os.Stdin)

		words := make(map[string]int)

		scan.Split(bufio.ScanWords)

		for scan.Scan() {
			words[scan.Text()]++
		}

		fmt.Println(len(words))
		fmt.Println(words)
	*/
	/*
		for _, fname := range os.Args[1:] {
			//fmt.Fprintln(os.Stdout, fname)
			file, err := os.Open(fname)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if _, err := io.Copy(os.Stdout, file); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			file.Close()
		}
	*/

	// Recursion from the GO Programming language book
	/*
			doc, err := html.Parse(os.Stdin)
			if err != nil {
				fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
				os.Exit(-1)
			}
			for _, link := range visit(nil, doc) {
				fmt.Println(link)
			}
			var mapping = make(map[string]int)
			summary(mapping, doc)
			//fmt.Println(mapping)
			for k, v := range mapping {
				fmt.Println(k, v)
			}
			text(doc)

			links, err := findlinks1("htthtps://www.google.co.in/")
			if err != nil {
				fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
				return
			}
			for _, link := range links {
				fmt.Println(link)
			}
			//s := outline(nil, doc)
			//fmt.Println(s)
		}

		func findlinks1(url string) ([]string, error) {
			resp, err := http.Get(url)
			if err != nil {
				return nil, err
			}
			if resp.StatusCode != http.StatusOK {
				resp.Body.Close()
				return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
			}
			doc, err := html.Parse(resp.Body)
			resp.Body.Close()
			if err != nil {
				return nil, fmt.Errorf("Parsing %s as HTML: %v", url, err)
			}
			return visit(nil, doc), nil
	*/
}

// ex 5.3 write a function to print the contents of all text nodes.
func text(doc *html.Node) {
	if doc.Type == html.ElementNode && (doc.Data == "style" || doc.Data == "script") {
		return
	}
	if doc.Type == html.TextNode {
		fmt.Println(doc.Data)
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		text(c)
	}
}

// ex 5.2
func summary(mapping map[string]int, doc *html.Node) {
	if doc.Type == html.ElementNode {
		mapping[doc.Data]++
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		summary(mapping, c)
	}
}

func outline(stack []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
	return stack
}

// ex 5.1
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	/*
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = visit(links, c)
		}
	*/
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}
