package gifenc

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"os"
)

func (config Config) Decode(path string) error {
	// Open the GIF
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error while opening file: %s", err)
	}

	split(file, config.Width, config.Height, config.Output)

	return nil
}

// Split the GIF into images
func split(file io.Reader, width int, height int, output Output) (err error) {
	defer func() {
		if recv := recover(); recv != nil {
			err = fmt.Errorf("error while decoding file: %s", recv)
		}
	}()

	gif, err := gif.DecodeAll(file)
	if err != nil {
		return fmt.Errorf("error while decoding file: %s", err)
	}

	x, y := getArea(gif)
	if width == 0 {
		width = x
	}
	if height == 0 {
		height = y
	}

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(dst, dst.Bounds(), gif.Image[0], image.Point{}, draw.Src)

	for i, img := range gif.Image {
		draw.Draw(dst, dst.Bounds(), img, image.Point{}, draw.Over)

		file, err := os.Create(fmt.Sprintf("%s%s%d%s", output.Path, output.Name, i, ".png"))
		if err != nil {
			return fmt.Errorf("error while creating image: %s", err)
		}

		err = png.Encode(file, dst)
		if err != nil {
			return fmt.Errorf("error while creating image: %s", err)
		}

		file.Close()
	}

	return nil
}

func getArea(gif *gif.GIF) (x, y int) {
	var xLow int
	var xHigh int
	var yLow int
	var yHigh int

	for _, img := range gif.Image {
		if img.Rect.Min.X < xLow {
			xLow = img.Rect.Min.X
		}
		if img.Rect.Min.Y < yLow {
			yLow = img.Rect.Min.Y
		}
		if img.Rect.Max.X > xHigh {
			xHigh = img.Rect.Max.X
		}
		if img.Rect.Max.Y > yHigh {
			yHigh = img.Rect.Max.Y
		}
	}

	return xHigh - xLow, yHigh - yLow
}
