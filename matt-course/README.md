package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("hash of %s is %s; length: %d", p.Path, p.Hash, p.Length)
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type Filenamer interface {
	Filename() string
}

func main() {
	p := Pair{"/usr", "0xfde"}
	var fn Filenamer = PairWithLength{Pair{"/usr/bin/", "0xabcq"}, 122}
	fmt.Println(fn)
	fmt.Println(p)
	fmt.Println(fn.Filename())
	fmt.Println(p.Filename())

}



package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

//func (s Organs) Less(i, j int) bool { return s[i].Name < s[j].Name }

type ByName struct{ Organs }
type ByWeight struct{ Organs }

func (s ByName) Less(i, j int) bool   { return s.Organs[i].Name < s.Organs[j].Name }
func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }

func main() {
	//s := Organs{{"brain", 120}, {"heart", 450}, {"liver", 1400}, {"pancreas", 320}}
	//sort.Sort(s)
	s := []Organ{{"brain", 120}, {"heart", 450}, {"liver", 1400}, {"pancreas", 320}}
	sort.Sort(ByName{s})
	fmt.Println(s)
	sort.Sort(ByWeight{s})
	fmt.Println(s)
}
