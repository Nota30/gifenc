package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"

	"github.com/Nota30/gifenc"
)

func main() {
	decode()
	encode()
}

// Encode
func encode() {
	init := gifenc.Config{
		Delay: 30,
	}

	images := getImages()

	encoded, err := init.Encode(images)
	if err != nil {
		fmt.Print(err)
	}

	newfile, err := os.Create(fmt.Sprintf("%s%s", "test/", "sword_test.gif"))
	if err != nil {
		fmt.Print(err)
	}
	defer newfile.Close()

	encodeErr := gif.EncodeAll(newfile, encoded)

	if encodeErr != nil {
		fmt.Print(err)
	}
}

// Decode
func decode() {
	init := gifenc.Config{
		Delay: 30,
	}

	file, err := os.Open("test/input/sword.gif")
	if err != nil {
		fmt.Print(err)
	}

	gif, err := gif.DecodeAll(file)
	if err != nil {
		fmt.Print(err)
	}

	imgs, err := init.Decode(gif)
	if err != nil {
		fmt.Print(err)
	}

	saveImages(imgs)
}

func saveImages(imgs []*image.RGBA) {
	for i, img := range imgs {
		file, err := os.Create(fmt.Sprintf("%s%s%d%s", "test/output/", "sword", i, ".png"))
		if err != nil {
			fmt.Print(err)
		}

		err = png.Encode(file, img)
		if err != nil {
			fmt.Print(err)
		}

		file.Close()
	}
}

// Get all the images inside a directory
func getImages() []image.Image {
	var images []image.Image
	files, err := os.ReadDir("test/output/")
	if err != nil {
		fmt.Print(err)
	}

	var allFiles []string
	for _, file := range files {
		allFiles = append(allFiles, file.Name())
	}

	for _, file := range allFiles {
		reader, err := os.Open("test/output/" + file)
		if err != nil {
			fmt.Print(err)
		}
		defer reader.Close()

		img, err := png.Decode(reader)
		if err != nil {
			fmt.Print(err)
		}

		images = append(images, img)
	}

	return images
}
