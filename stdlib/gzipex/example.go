package gzipex

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CompresFile() {
	// Some text we want to compress.
	original := "bird and frog"

	// Open a file for writing.
	f, err := os.Create("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/gzipex/test.gz")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Create gzip writer.
	w := gzip.NewWriter(f)

	// Write bytes in compressed form to the file.
	w.Write([]byte(original))

	// Close the file.
	w.Close()

	fmt.Println("DONE")
}

// Decompress. Here we decompress data from a file on the disk. We first read in the file with os.Open.
// Then we create a gzip Reader with the gzip.NewReader call.
func DecompressGZ() {
	// Open the gzip file.
	f, _ := os.Open("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/gzipex/test.gz")

	// Create new reader to decompress gzip.
	reader, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Empty byte slice.
	result := make([]byte, 100)

	// Read in data.
	count, _ := reader.Read(result)

	// Print our decompressed data.
	fmt.Println(count)
	fmt.Println(string(result))
}

// BestCompression. We can adjust the GZIP algorithm to emphasize speed or the output's
//compressed file size. Here we build up a string that contains enough characters to show a difference.

func BestCompression() {
	test := "this is an example string for testing compression levels"
	test += " here is some more example text"
	test += " cat 123 bird 456 UPPERCASE TEXT"
	test += " __ punct __ // punct"

	// Write with BestSpeed.
	fmt.Println("BESTSPEED")
	f, _ := os.Create("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/gzipex/bestspeed.gz")
	w, _ := gzip.NewWriterLevel(f, gzip.BestSpeed)
	w.Write([]byte(test))
	w.Close()

	// Write with BestCompression.
	fmt.Println("BESTCOMPRESSION")
	f, _ = os.Create("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/gzipex/bestcompression.gz")
	w, _ = gzip.NewWriterLevel(f, gzip.BestCompression)
	w.Write([]byte(test))
	w.Close()
}

// Compress file. This program reads in a file. It uses ioutil.ReadAll
// to get all the bytes from the file. It then creates a new file by replacing the extension with "gz" and writing it.
func CompressFile() {
	// Open file on disk.
	name := "test.txt"
	f, _ := os.Open("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/gzipex/" + name)

	// Create a Reader and use ReadAll to get all the bytes from the file.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Replace txt extension with gz extension.
	name = strings.Replace(name, ".txt", ".gz", -1)

	// Open file for writing.
	f, _ = os.Create("/Users/karel/Desktop/projeckty/gostd_examples/stdlib/gzipex/" + name)

	// Write compressed data.
	w := gzip.NewWriter(f)
	w.Write(content)
	w.Close()

	// Done.
	fmt.Println("DONE")
}
