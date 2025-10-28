package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	src := dir + "/file-handle/experiment/file1.txt"
	dst := dir + "/file-handle/experiment/file2.txt"

	srcData, _ := os.ReadFile(src)
	err := os.WriteFile(dst, srcData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
