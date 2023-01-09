package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func generateBlackSquareImage(size int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	return image.NewPaletted(image.Rect(0, 0, size, size), palette)
}

func generateSineGif() gif.GIF {
	var images []*image.Paletted
	var delays []int

	image := generateBlackSquareImage(1000)
	images = append(images, image)
	delays = append(delays, 0)

	return gif.GIF{
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

func main() {
	fmt.Println("Hello Sine Wave Exercise")

	image := generateSineGif()
	saveGifToFile(image, "sinWave.gif")
}