package bytesexample

import (
	"bytes"
	"fmt"
)

func BytesContains() {
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte("sea food"), []byte(" ")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))
}

// kontroluje jestli předaný []byte("I like seafood.") obsahuje něco z fÄo!
func BytesContainsAny() {
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!"))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的."))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), ""))
	fmt.Println(bytes.ContainsAny([]byte(""), ""))
}

func BytesContainsRune() {
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 's'))
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'ö'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '!'))
	fmt.Println(bytes.ContainsRune([]byte(""), '@'))
}

func BytesCount() {
	// kolik obsahuje cheese písmen e
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune -> runa je každé písmeno
}

func BytesCut() {
	// definice funkcionality pro check nad stringem
	show := func(s, sep string) {
		before, after, found := bytes.Cut([]byte(s), []byte(sep))
		// found -> vrací true/false podle toho jestli ve strungu "ph" našel
		// before= to co je v zadaném stringu Gopher před stringem ph
		// after=  to co je v zadaném stringu za za stringem "ph" podle kterého pracuju
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "ph")
}

func BytesEqual() {
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))
}

func BytesFields() {
	// rozdělí předanej []byte na [][]string takže na každém indexu je jedno slovo
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("foo bar baz taz")))
}

// kontroluje jestli předaný slice byte začíná předaným slice byte
func BytesHasPrefix() {
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))
}
