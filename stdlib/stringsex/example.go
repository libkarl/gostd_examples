package stringsex

import "fmt"

func SubStringEx1() {
	v := "value"
	subv := v[1:5]
	fmt.Println(subv)
}

func SubStringEx2() {
	// This string contains Unicode characters.
	value := "ü:ü"
	// Convert string to rune slice before taking substrings.
	runes := []rune(value)
	fmt.Println("First 1:", string(runes[0:1]))
	fmt.Println(" Last 1:", string(runes[2:]))
}
