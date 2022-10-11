package regexpex

import (
	"fmt"
	"regexp"
)

func MatchStringEx() {
	value := "gopher"

	// See if this regular expression matches.
	matched, _ := regexp.MatchString(".+r", value)

	// Test the result.
	if matched {
		fmt.Println(true)
	}
}

// Compile regexp. A regexp can be compiled.
//Once compiled, it can be reused many times with no delays before calls. Many methods, like MatchString are available on regexp instances.

func CompileRegex() {
	// Compile this regular expression.
	var lettersDigits = regexp.MustCompile(`\w\w\w\d\d\d`)

	// Some strings to test.
	values := []string{"cat100", "dog200",
		"bird", "200fish", "500bear", "frog800"}

	// Loop over our strings and call MatchString on them.
	for i := range values {
		result := lettersDigits.MatchString(values[i])
		fmt.Printf("%v = %v\n", values[i], result)
	}
}
