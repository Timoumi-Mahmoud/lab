package main

import (
	"fmt"
	"path/filepath"
	"sort"
)

type Pair struct {
	Path string
	Hash string
}
type PairWithLength struct {
	Pair   //embedded
	Length int
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %v is %v", p.Path, p.Hash)
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Length of %v is %v with hash %v",
		p.Path, p.Length, p.Hash)
}

/*
func Filename(p Pair) string {
	return filepath.Base(p.Path)
}
*/

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

func main() {
	pl := PairWithLength{Pair{"/usr", "0xfdfe"}, 121}
	fmt.Println("-:", pl.Path, pl.Hash, pl.Length) // not pl.x.Path
	fmt.Println(pl.String())
	fmt.Println(pl)
	fmt.Println()
	p := Pair{"/etc/vim", "0xdead"}
	/*fmt.Println(Filename(p))
	fmt.Println(Filename(pl.Pair))
	*/
	fmt.Println(pl.Filename())
	fmt.Println(p.Filename())

	fmt.Println()
	var fn Filenamer = PairWithLength{Pair{"/boot", "0xfuckoff"}, 666}

	fmt.Println(fn.Filename())
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	entries := []string{"charlie", "able", "dog", "baker"}
	sort.Sort(sort.StringSlice(entries))
	fmt.Println(entries)
	sort.Sort(sort.Reverse(sort.StringSlice(entries)))
	fmt.Println(entries)

}

type Filenamer interface {
	Filename() string
}

/*

struct composition : build struct from other struct
composition not inheritance

*/
