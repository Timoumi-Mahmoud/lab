package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	p := filepath.Join("dir1", "dir2", "dir3")

	fmt.Println(filepath.Join("dir1//", "filename"))

	fmt.Println("p:", p)
	fmt.Println("Dir(p): ", filepath.Base(p))
	filename := "img.jpg"
	ext := filepath.Ext(filename)
	fmt.Println("ext: ", ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}

	fmt.Println(rel)

}
