package main

import (
	"fmt"

	"github.com/Nota30/gifenc"
)

func main() {
	init := gifenc.Config{
		Output: gifenc.Output{
			Name: "sword1",
			Path: "test/output/",
		},
		Delay: 20,
	}

	err := init.Decode("test/input/sword.gif")
	if err != nil {
		fmt.Print(err)
	}

	// err := init.Encode("test/output/")
	// if err != nil {
	// 	fmt.Print(err)
	// }
}
