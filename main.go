package main

import (
	"fmt"
	"path"
)

func main() {
	var file, dir string
	dir, file = path.Split("css/main.css")
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}
