package pictures

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func PictureExample1() {
	catFile, err := os.Open("/Users/karel/projects/gostd_examples/stdlib/pictures/folder.png")
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	imData, imType, err := image.Decode(catFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(imData)
	fmt.Println(imType)

	cat, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cat)
}

func PictureExampleTerminalPrint() {
	// This example uses png.Decode which can only decode PNG images.
	catFile, err := os.Open("/Users/karel/projects/gostd_examples/stdlib/pictures/folder.png")
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	// Consider using the general image.Decode as it can sniff and decode any registered image format.
	img, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}
	// vypíše číselnou reprezentaci RGB v matrixu, který je výsledkem png.Decode
	fmt.Println(img)

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}
