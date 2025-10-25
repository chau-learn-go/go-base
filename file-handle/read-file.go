package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	//f, err := os.Open(dir + "/file-handle/experiment/file1.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer f.Close()
	//buf := make([]byte, 100)
	n, _ := os.ReadFile(dir + "/file-handle/experiment/file1.txt")
	fmt.Println(string(n))
}
