package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

func generateBlackSquareImage(size int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	return image.NewPaletted(image.Rect(0, 0, size, size), palette)
}

func generateSineGif(imageCh chan gif.GIF) {
	var images []*image.Paletted
	var delays []int
	size := 1000
	frequencyCoefficient := 16.0
	amplitudeCoefficient := 50.0

	image := generateBlackSquareImage(size)
	for x := 0; x < size; x++ {
		sine := math.Sin(float64(x)/frequencyCoefficient) * amplitudeCoefficient
		y := int(sine) + (size / 2)
		image.SetColorIndex(x, y, 1)
	}

	images = append(images, image)
	delays = append(delays, 0)

	imageCh <- gif.GIF{
		Image: images,
		Delay: delays,
	}
}

func saveGifToFile(image gif.GIF, fileName string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &image)
}

func getFileName(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}

func main() {
	fmt.Println("Hello Sine Wave Exercise")

	ch := make(chan string)
	go getFileName(ch)
	var fileName string

	imageCh := make(chan gif.GIF)
	go generateSineGif(imageCh)
	var image gif.GIF

	for i := 0; i < 2; i++ {
		select {
		case word := <-ch:
			fileName = word
		case img := <-imageCh:
			image = img
		}
	}

	saveGifToFile(image, fileName)
}
