package main

import "fmt"

const A int = 1

const B float32 = 3.14

const C string = "constant"

const D byte = 'D'

func main() {
	fmt.Printf("%d, %f, %s, %d", A, B, C, D)
}
