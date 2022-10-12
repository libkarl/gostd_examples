package filework

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func FileWorkExample() {
	// Open the file.
	f, _ := os.Open("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/test.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func ScanWords() {
	f, _ := os.Open("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/test.txt")
	scanner := bufio.NewScanner(f)

	// Set the Split method to ScanWords.
	scanner.Split(bufio.ScanWords)

	// Scan all words from the file.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

// Read entire file. With ioutil.ReadAll we can get the entire contents of a file in a byte slice or a string.
// We must import "io/ioutil" and then call ioutil.ReadAll on a reader.
func ReadAll() {
	// Open a file.
	f, _ := os.Open("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/test.txt")

	// Use bufio.NewReader to get a Reader.
	// ... Then use ioutil.ReadAll to read the entire content.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// File content.
	fmt.Println(string(content))
}

func FileExists() {
	directory := "/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/"
	_, err := os.Stat(directory)

	// See if directory exists.
	// ... Use the IsNotExist method.
	if os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
	}

	file := "/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/test.txt"
	_, err = os.Stat(file)

	// See if the file exists.
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	}
}

// Get file size. Sometimes we wish to get the size in bytes of a file. We can call os.Stat and then the Size() method. This returns the count of bytes in the file.
func FileSize() {
	file := "/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/test.txt"

	// Call Stat on a path string to get statistics.
	stat, err := os.Stat(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Get file size.
	size := stat.Size()
	fmt.Println("FILE SIZE IN BYTES:", size)
}

// Move file. How can we rename or move a file? We must import the "os" package,
// which contains many helpful file-system funcs. In main(), we have "before" and "after" locations.
func MoveFile() {

	before := "/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/movedir/test.txt"
	after := "/Users/karel/Desktop/projeckty/gostd_examples/stdlib/filework/test.txt"

	// Rename or move file from one location to another.
	os.Rename(before, after)
	fmt.Println("DONE")
}

func SeekerTest() {
	file, _ := os.Open("/Users/karel/projects/gostd_examples/stdlib/filework/test.txt")
	reader := io.Reader(file)
	buffer, err := io.ReadAll(reader)
	fmt.Printf("ReadAll buffer={%v}, err={%v}\n", string(buffer), err)
	seeker := reader.(io.Seeker)
	seeker.Seek(0, io.SeekStart)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("ReadAll buffer={%v}, err={%v}\n", string(buffer), err)
	file.Close()
}
