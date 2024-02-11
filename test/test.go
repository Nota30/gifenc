package main

import (
	"fmt"
	"image/gif"
	"image/png"
	"os"

	"github.com/Nota30/gifenc"
)

func main() {
	init := gifenc.Config{
		Delay: 30,
	}

	// DECODE
	imgs, err := init.Decode("test/input/sword.gif")
	if err != nil {
		fmt.Print(err)
	}

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

	// ENCODE
	encoded, err := init.Encode("test/output/")
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
