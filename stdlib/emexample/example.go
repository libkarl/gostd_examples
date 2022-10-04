package emexample

import (
	"embed"
	"fmt"
)

var f embed.FS

func EmbedFile() {
	//go:embed hello.txt
	data, err := f.ReadFile("hello.txt")
	if err != nil {
		fmt.Print("read file fuckup")
	}
	fmt.Print(string(data))
}
