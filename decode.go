package gifenc

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
)

func (config Config) Decode(gif *gif.GIF) ([]*image.RGBA, error) {
	imgs, err := split(gif, config.Width, config.Height)
	if err != nil {
		return nil, fmt.Errorf("error while decoding file: %s", err)
	}

	return imgs, nil
}

// Split the GIF into images
func split(gif *gif.GIF, width int, height int) (imgs []*image.RGBA, err error) {
	defer func() {
		if recv := recover(); recv != nil {
			err = fmt.Errorf("error while decoding file: %s", recv)
		}
	}()

	x, y := getArea(gif)
	if width == 0 {
		width = x
	}
	if height == 0 {
		height = y
	}

	var images []*image.RGBA
	for _, img := range gif.Image {
		dst := image.NewRGBA(image.Rect(0, 0, width, height))
		draw.Draw(dst, dst.Bounds(), img, image.Point{}, draw.Over)

		new := dst
		images = append(images, new)
	}

	return images, nil
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
