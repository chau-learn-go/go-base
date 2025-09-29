package main

import "fmt"

func main() {
	var m1 map[string]int                // nil map
	m2 := make(map[string]int)           // usable
	m3 := map[string]int{"a": 1, "b": 2} // literal

	fmt.Println("Bài 1:")
	fmt.Println("m1:", m1) // nil
	fmt.Println("m2:", m2)
	fmt.Println("m3:", m3)

	fmt.Println()

	ages := make(map[string]int)
	ages["Alice"] = 20
	ages["Bob"] = 25

	fmt.Println("Bài 2:")
	fmt.Println("ages:", ages)
	fmt.Println("Alice:", ages["Alice"])

	v, ok := ages["Charlie"]
	fmt.Println("Charlie exists?", ok, "value:", v)
	fmt.Println()

	delete(ages, "Bob")
	fmt.Println("Bài 3:")
	fmt.Println("ages sau khi xóa Bob:", ages)
	fmt.Println()

	fmt.Println("Bài 4:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
	fmt.Println()

	class := make(map[string]map[string]int)
	class["10A"] = map[string]int{"An": 8, "Bình": 9}
	class["10B"] = map[string]int{"Cường": 7}

	fmt.Println("Bài 5:")
	for className, students := range class {
		fmt.Println("Lớp:", className)
		for student, score := range students {
			fmt.Printf("  %s: %d\n", student, score)
		}
	}
	fmt.Println()

	countries := make(map[string][]string)
	countries["VN"] = []string{"Hanoi", "Saigon"}
	countries["US"] = []string{"New York"}

	// Append thêm city
	countries["VN"] = append(countries["VN"], "Danang")

	fmt.Println("Bài 6:")
	for country, cities := range countries {
		fmt.Println(country, ":", cities)
	}
	fmt.Println()
}
