# Gifenc

A Golang based GIF encoder/decoder
- Convert GIFs to images and images to GIFs

## Prerequisites
- [Go](https://go.dev/)
 
## Quick Start
Import the `gifenc` package first.
```go
package main

import (
	"github.com/Nota30/gifenc"
)

func main() {
	gify := gifenc.Config{}

    // Decode
	var gif *gif.GIF // Provide a gif
	imgs, err := gify.Decode(gif)
	if err != nil {
		println(err)
	}

    // Encode
	var images []image.Image // Provide an array of images
    encoded, err := gify.Encode(images)
    if err != nil {
		println(err)
	}
}
```
You can view examples in the `test/test.go` file.

## Issues
- This package uses Floyd–Steinberg dithering so the GIF result might not be what you expected.
- Maybe adding other dithering algorithms at a later release?
