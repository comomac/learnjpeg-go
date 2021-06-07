package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func saveJPG() {
	fileIn := "sample-32x24-p6.ppm"
	fileOut := "sample-32x24-p6-go.jpg"

	dat, err := ioutil.ReadFile(fileIn)
	if err != nil {
		panic(err)
	}

	fmt.Printf("filesize: %d bytes\n", len(dat))

	colours, img, err := PPMDecode(dat)
	if err != nil {
		panic(err)
	}

	fmt.Printf("image w x h: %d x %d\n", img.Rect.Dx(), img.Rect.Dy())
	fmt.Printf("colour max:  %d\n", colours)

	out, err := os.Create(fileOut)
	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("image converted")
}
