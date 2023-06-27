package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWords(s string) int {
	words := strings.Split(s, " ")
	count := 0
	for _, word := range words {
		word = strings.TrimRight(word, ".,?!")
		word = strings.Replace(word, "-", "", -1)
		if validWord(word) {
			count++
		}
	}
	return count
}

func validWord(word string) bool {
	re := regexp.MustCompile(`[^a-zA-Z\s]`)
	return !re.MatchString(word)
}

func main() {
	case1 := "Saat meng*ecat tembok, Agung dib_antu oleh Raihan."
	fmt.Println("Hasil:", countWords(case1))

	case2 := "Berapa u(mur minimal[ untuk !mengurus ktp?"
	fmt.Println("Hasil:", countWords(case2))

	case3 := "Masing-masing anak mendap(atkan uang jajan ya=ng be&rbeda."
	fmt.Println("Hasil:", countWords(case3))
}
