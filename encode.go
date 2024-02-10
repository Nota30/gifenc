package gifenc

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
)

func (config Config) Encode(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error while opening dir: %s", err)
	}

	var allFiles []string
	for _, file := range files {
		fmt.Print(file.Name())
		allFiles = append(allFiles, file.Name())
	}

	animated := gif.GIF{
		LoopCount: len(allFiles),
	}

	for _, file := range allFiles {
		reader, err := os.Open(path + file)
		if err != nil {
			return fmt.Errorf("error while opening file: %s", err)
		}
		defer reader.Close()

		img, err := png.Decode(reader)
		if err != nil {
			return fmt.Errorf("error while decoding image: %s", err)
		}
		bounds := img.Bounds()
		drawer := draw.FloydSteinberg

		paletted := image.NewPaletted(bounds, palette.Plan9)

		drawer.Draw(paletted, img.Bounds(), img, image.Point{})
		animated.Image = append(animated.Image, paletted)
		animated.Delay = append(animated.Delay, config.Delay)
	}

	file, err := os.Create(config.Output.Name)
	if err != nil {
		return fmt.Errorf("error while creating file: %s", err)
	}
	defer file.Close()

	encodeErr := gif.EncodeAll(file, &animated)
	if encodeErr != nil {
		return fmt.Errorf("error while creating file: %s", err)
	}

	return nil
}
