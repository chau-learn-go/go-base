package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]string{"Go", "Rust", "C++"}

	fmt.Println("Bài 1:")
	fmt.Println("arr1:", arr1, "len:", len(arr1))
	fmt.Println("arr2:", arr2, "len:", len(arr2))
	fmt.Println("arr3:", arr3, "len:", len(arr3))

	var arr4 [5]int
	arr4[0] = 10
	arr4[1] = 20
	arr4[2] += 10
	fmt.Println("\nBài 2:")
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("arr4[%d] = %d\n", i, arr4[i])
	}

	a := [3]int{1, 2, 3}
	b := a // copy by value
	b[0] = 100
	fmt.Println("\nBài 3:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a2 := [2]int{1, 2}
	b2 := [2]int{1, 2}
	fmt.Println("\nBài 4:")
	fmt.Println("a2 == b2 ?", a2 == b2)
	b2[1] = 3
	fmt.Println("a2 == b2 ?", a2 == b2)

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	fmt.Println("\nBài 5:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
