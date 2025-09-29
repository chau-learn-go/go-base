package main

import "fmt"

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

func main() {

	var s1 []int
	s2 := []string{"Go", "Rust", "C++"}
	s3 := make([]int, 5)

	fmt.Println("Bài 1:")
	fmt.Println("s1:", s1, "len:", len(s1), "cap:", cap(s1))
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2))
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))
	fmt.Println()

	arr := [6]int{1, 2, 3, 4, 5, 6}
	sa := arr[1:4]
	sb := arr[:3]
	sc := arr[3:]
	fmt.Println("Bài 2:")
	fmt.Println("arr:", arr)
	fmt.Println("sa:", sa)
	fmt.Println("sb:", sb)
	fmt.Println("sc:", sc)
	sa[0] = 100
	fmt.Println("sau khi thay đổi sa[0] = 100 → arr:", arr)
	fmt.Println()

	sl := []int{1, 2, 3}
	fmt.Println("Bài 3:")
	fmt.Println("Trước append:", sl, "len:", len(sl), "cap:", cap(sl))
	sl = append(sl, 4, 5, 6)
	fmt.Println("Sau append thêm 4,5,6:", sl, "len:", len(sl), "cap:", cap(sl))
	sl = append(sl, []int{7, 8}...)
	fmt.Println("Sau append thêm slice {7,8}:", sl, "len:", len(sl), "cap:", cap(sl))
	fmt.Println()

	src := []int{1, 2, 3, 4}
	dst := make([]int, 2)
	copy(dst, src)
	fmt.Println("Bài 4:")
	fmt.Println("src:", src)
	fmt.Println("dst (sau copy):", dst)
	fmt.Println()

	matrix := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	matrix = append(matrix, row1, row2)
	fmt.Println("Bài 6:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()

	sl2 := []int{10, 20, 30}
	fmt.Println("Bài 7:")
	fmt.Println("Trước modify:", sl2)
	modifySlice(sl2)
	fmt.Println("Sau modify:", sl2)
	fmt.Println()
}
