package main

import "fmt"

func main() {
	s1 := "Hello"
	var s2 string
	s3 := `Line1 Line2 Line3`

	fmt.Println("Bài 1:")
	fmt.Println("s1:", s1, "len:", len(s1))
	fmt.Println("s2:", s2, "len:", len(s2))
	fmt.Println("s3:", s3, "len:", len(s3))
	fmt.Println()

	s := "Xin chào"
	fmt.Println("Bài 2:")
	fmt.Println("Duyệt bằng byte (for i):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %v (%c)\n", i, s[i], s[i])
	}

	fmt.Println("Duyệt bằng rune (for range):")
	for i, r := range s {
		fmt.Printf("s[%d] = %U (%c)\n", i, r, r)
	}
	fmt.Println()

	part1 := "Go"
	part2 := "lang"
	full := part1 + part2
	full2 := fmt.Sprintf("%s%s", part1, part2)
	fmt.Println("Bài 3:")
	fmt.Println("full:", full)
	fmt.Println("full2:", full2)
	fmt.Println()

	word := "Golang"
	fmt.Println("Bài 4:")
	fmt.Println("word[0:2] =", word[0:2]) // "Go"
	fmt.Println("word[2:5] =", word[2:5]) // "lan"
	fmt.Println("word[3:]  =", word[3:])  // "ang"
	fmt.Println()

	fmt.Println("Bài 5:")
	fmt.Println(`"apple" < "banana"?`, "apple" < "banana")
	fmt.Println(`"go" == "Go"?`, "go" == "Go")
	fmt.Println()

	unistr := "Xin chào"
	fmt.Println("Bài 6:")
	fmt.Println("len(unistr) =", len(unistr))                 // số byte
	fmt.Println("len([]rune(unistr)) =", len([]rune(unistr))) // số ký tự thật sự
	fmt.Println("[]byte(unistr):", []byte(unistr))
	fmt.Println("[]rune(unistr):", []rune(unistr))
	fmt.Println()

	sImm := "hello"
	// sImm[0] = 'j' // lỗi compile
	runes := []rune(sImm)
	runes[0] = 'j'
	newStr := string(runes)
	fmt.Println("Bài 7:")
	fmt.Println("Trước:", sImm)
	fmt.Println("Sau:", newStr)
	fmt.Println()
}
