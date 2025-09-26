package main

import "fmt"

type Point struct {
	X, Y int
}

type Greeter interface {
	Greet() string
}

var b bool

var i int
var f float64
var c complex128

var s string

var arr [3]int

var sl []int

var p Point

var ptr *Point

var fn func(a, b int) int

var g Greeter

var m map[string]int

var ch chan int

var x interface{}

func main() {
	k := 1
	fmt.Println(k)
}
