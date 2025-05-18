package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func add2(a, b int) int {
// 	return a + b
// }

func getLanguages() (string, string, int) {
	return "java", "javascript", 2
}

func main() {
	// fmt.Println(add(2, 5))
	lang, lang2, _ := getLanguages()
	fmt.Println(lang, lang2)
}
