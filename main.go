package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// readPPM()
	a := jheader(1, 2)
	fmt.Printf("%d  %+v\n", len(a), a)
	//saveJPG()
}

func readPPM() {
	file := "sample-256x192-p6.ppm"

	dat, err := ioutil.ReadFile(file)
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
}
