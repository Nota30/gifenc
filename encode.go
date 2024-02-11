package gifenc

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
)

// Encode a GIF. Provide an array of image to be combined into a GIF.
func (config Config) Encode(images []image.Image) (*gif.GIF, error) {
	animated := gif.GIF{
		LoopCount: 0,
	}

	for _, img := range images {
		bounds := img.Bounds()
		drawer := draw.FloydSteinberg

		paletted := image.NewPaletted(bounds, palette.Plan9)

		drawer.Draw(paletted, img.Bounds(), img, image.Point{})
		animated.Image = append(animated.Image, paletted)
		animated.Delay = append(animated.Delay, config.Delay)
	}

	return &animated, nil
}
