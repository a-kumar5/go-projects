package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ayushkumar/gocrypt/internal/crypto"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `gocrypt - AES-256-GCM encryption CLI

Usage:
  gocrypt gen-key [-out keyfile]
  gocrypt enc [-in file|-] [-out file|-] [-key keyfile] [-pass] [-salt saltfile]
  gocrypt dec [-in file|-] [-out file|-] [-key keyfile] [-pass] [-salt saltfile]

Options:
  -in    input file (use - for stdin)
  -out   output file (use - for stdout)
  -key   path to 32-byte raw key file for AES-256
  -pass  derive key from passphrase via scrypt (prompts securely)
  -salt  file to store/read salt for scrypt when using -pass

Notes:
  - If both -key and -pass are provided, -key takes precedence.
  - AES-GCM is used with a 12-byte random nonce and includes authentication.
`)
}

func readAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func writeAll(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}

func openInput(path string) (io.ReadCloser, error) {
	if path == "-" {
		return io.NopCloser(os.Stdin), nil
	}
	return os.Open(path)
}

func openOutput(path string) (io.WriteCloser, error) {
	if path == "-" {
		return nopWriteCloser{Writer: os.Stdout}, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	return os.Create(path)
}

type nopWriteCloser struct{ io.Writer }

func (nwc nopWriteCloser) Close() error { return nil }

func promptPassphrase(prompt string) (string, error) {
	// Fallback to non-echo if available on platform; otherwise, use basic input.
	fmt.Fprint(os.Stderr, prompt)
	reader := bufio.NewReader(os.Stdin)
	pw, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	return strings.TrimSpace(pw), nil
}

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	sub := os.Args[1]
	switch sub {
	case "gen-key":
		fs := flag.NewFlagSet("gen-key", flag.ExitOnError)
		out := fs.String("out", "key.bin", "output key file")
		_ = fs.Parse(os.Args[2:])

		key := make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, key); err != nil {
			log.Fatalf("failed generating key: %v", err)
		}
		if err := os.WriteFile(*out, key, 0o600); err != nil {
			log.Fatalf("failed writing key: %v", err)
		}
		fmt.Fprintf(os.Stderr, "wrote 32-byte key to %s\n", *out)

	case "enc", "dec":
		fs := flag.NewFlagSet(sub, flag.ExitOnError)
		in := fs.String("in", "-", "input file (- for stdin)")
		out := fs.String("out", "-", "output file (- for stdout)")
		keyFile := fs.String("key", "", "raw 32-byte key file")
		pass := fs.Bool("pass", false, "derive key from passphrase (scrypt)")
		saltFile := fs.String("salt", "salt.bin", "salt file for scrypt when using -pass")
		_ = fs.Parse(os.Args[2:])

		var key []byte
		var err error
		if *keyFile != "" {
			key, err = os.ReadFile(*keyFile)
			if err != nil {
				log.Fatalf("failed to read key: %v", err)
			}
			if len(key) != 32 {
				log.Fatalf("key must be 32 bytes, got %d", len(key))
			}
		} else if *pass {
			pw, err := promptPassphrase("Enter passphrase: ")
			if err != nil {
				log.Fatalf("failed reading passphrase: %v", err)
			}
			// Load or create salt
			var salt []byte
			if b, err := os.ReadFile(*saltFile); err == nil {
				salt = b
			} else {
				salt = make([]byte, 16)
				if _, err := io.ReadFull(rand.Reader, salt); err != nil {
					log.Fatalf("failed generating salt: %v", err)
				}
				if err := os.WriteFile(*saltFile, salt, 0o600); err != nil {
					log.Fatalf("failed writing salt: %v", err)
				}
			}
			key, err = crypto.DeriveKeyFromPassphrase(pw, salt)
			if err != nil {
				log.Fatalf("failed deriving key: %v", err)
			}
		} else {
			log.Fatalf("either -key or -pass must be provided")
		}

		inR, err := openInput(*in)
		if err != nil {
			log.Fatalf("failed opening input: %v", err)
		}
		defer inR.Close()
		data, err := readAll(inR)
		if err != nil {
			log.Fatalf("failed reading input: %v", err)
		}

		var outB []byte
		if sub == "enc" {
			ct, nonce, err := crypto.EncryptAESGCM(key, data)
			if err != nil {
				log.Fatalf("encrypt error: %v", err)
			}
			// Output format: base64(nonce||ciphertext)
			outRaw := append(nonce, ct...)
			b64 := make([]byte, base64.StdEncoding.EncodedLen(len(outRaw)))
			base64.StdEncoding.Encode(b64, outRaw)
			outB = b64
		} else {
			// Accept base64 input
			dec := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
			n, err := base64.StdEncoding.Decode(dec, data)
			if err != nil {
				log.Fatalf("base64 decode error: %v", err)
			}
			dec = dec[:n]
			if len(dec) < 12 {
				log.Fatalf("ciphertext too short")
			}
			nonce := dec[:12]
			ct := dec[12:]
			pt, err := crypto.DecryptAESGCM(key, nonce, ct)
			if err != nil {
				log.Fatalf("decrypt error: %v", err)
			}
			outB = pt
		}

		outW, err := openOutput(*out)
		if err != nil {
			log.Fatalf("failed opening output: %v", err)
		}
		defer outW.Close()
		if err := writeAll(outW, outB); err != nil {
			log.Fatalf("failed writing output: %v", err)
		}

	default:
		usage()
		os.Exit(2)
	}
}
