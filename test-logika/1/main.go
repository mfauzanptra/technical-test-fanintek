package main

import (
	"fmt"
)

func countSocks(socks []int) int {
	total := 0
	countSocks := map[int]int{}

	for _, sock := range socks {
		countSocks[sock]++
	}

	for _, count := range countSocks {
		total += count / 2
	}

	return total
}

func main() {
	case1 := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println("Hasil: ", countSocks(case1))

	case2 := []int{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5}
	fmt.Println("Hasil: ", countSocks(case2))

	case3 := []int{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	fmt.Println("Hasil: ", countSocks(case3))
}
