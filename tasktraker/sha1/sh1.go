package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf(sig)
}

// $ cat http.log.gz | gunzip | sha1sum
func sha1Sum(fileName string) (string, error) {
	// idiom: acquire a resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}

	defer file.Close() // deferred are called in LIFO order

	r, err := gzip.NewReader(file)
	//io.CopyN(os.Stdout, r, 100)
	if err != nil {
		return "", err
	}
	defer r.Close()

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
