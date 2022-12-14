package bufioexample

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Bufioexample() {
	// An artificial input source.
	const input = "1234 5678 12345678910 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

func EmptyFinalToken() {
	// Comma-separated list; last entry is empty.
	const input = "1,2,3,4,5"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// definice toho čeho se chci ve stringu zbavova
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	// Split dostane definici toho na čem má split provádět, vždy když dojde na znak , provede split
	scanner.Split(onComma)
	// proscanuje scanner a pomocí scanner.Text se vždy vypíše aktuální token na kterém for loop zrovna je
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func Words() {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York on Monday.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation. ScanWords is a split function for a Scanner
	// that returns each space-separated word of text, with surrounding spaces deleted.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	// ve scanneru jsou všechna slova ze stringu zbavená mezer díky pře definované funkci ScanWords
	// ve for loop je spočítá a vypíše jejich počet do terminálu
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}

// spočítá bytes načteného slova pomocí scanneru
func Bytes() {
	scanner := bufio.NewScanner(strings.NewReader("kokotihlava"))
	for scanner.Scan() {
		// pokud se počet písmen slova načteného do scanneru == 6 vypíše true
		fmt.Println(len(scanner.Bytes()) == 11)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}

func BufioWritter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "fucking ")
	fmt.Fprint(w, "world!\n")
	// aby writter z buffio vypsal do konzole co se do něj naposílalo je nutné nad ním zavolat w.Flush()
	w.Flush() // Don't forget to flush!
	h := bufio.NewWriter(os.Stdout)
	fmt.Fprint(h, "I am ")
	fmt.Fprint(h, "here!")
	h.Flush()
}

func AvailableBuffer() {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		// do b předá při každé iteraci string formu přečteného int64
		b = strconv.AppendInt(b, i, 10)

		// mezi každé číslo přidá do vypsaného stringu mezeru
		b = append(b, ' ')
		// zapíše obsah b do w
		w.Write(b)
	}
	w.Flush()
}

// 1
func BufferExample1() {
	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Ranjan")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is", strBuffer.String())
}

func BufferExample2() {
	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	len, _ := fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")

	fmt.Println("The length of given string is: ", len)
	strBuffer.WriteString("[DONE]")
	fmt.Println("The string buffer output is", strBuffer.String())
}

func BufferExample3() {
	//Creating buffer variable to hold and manage the string data
	var byteString bytes.Buffer
	// bytes package má i svou vlastní metodu pro plnění bufferu
	byteString.Write([]byte("Hello "))
	fmt.Fprintf(&byteString, "Hello friends how are you")
	// stejně tak má svou vlastní metodu pro zapisování s cílem
	var buffer2 bytes.Buffer
	byteString.WriteTo(&buffer2)

	// zapisuje do standard output
	buffer2.WriteTo(os.Stdout)
}

func BufferExample4() {
	//Creating buffer variable to hold and manage the string data
	var strByyte bytes.Buffer
	// předchystá kapacitu bufferu, aby nemuselo docházet k alokování paměti při zápisu
	strByyte.Grow(64)
	// zapíše do bufferu data
	strByyte.Write([]byte("It is a 64 byte"))
	// přeuloží bytes uložené v bufferu do nové proměnné
	strByytestrByyte := strByyte.Bytes()
	fmt.Println(string(strByytestrByyte[:]))
	fmt.Printf("%b", strByytestrByyte[:strByyte.Len()])
}

// buffion exampkle 5
type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("%q\n", p)
	return len(p), nil
}
func BufferExample5() {
	//
	s := strings.NewReader("onetwothreeegap")
	// vrací pointer na nově alokovanou strukturu v paměti
	w := new(Writer)
	// jako argument dám metodu, která je v souladu  interfacem a stanovím velikost
	bw := bufio.NewWriterSize(w, 3)
	bw.ReadFrom(s)
	err := bw.Flush()
	if err != nil {
		panic(err)
	}
}

//
