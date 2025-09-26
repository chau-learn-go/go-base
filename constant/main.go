package main

import "fmt"

const A int = 1

const B float32 = 3.14

const C string = "constant"

const D byte = 'D'

const (
	Sun = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

func main() {
	fmt.Printf("%d, %f, %s, %d\n", A, B, C, D)
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d", Sun, Mon, Tue, Wed, Thu, Fri, Sat)
}
