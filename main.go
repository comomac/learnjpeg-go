package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	file := "sample-256x192-p6.ppm"

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("filesize: %d bytes\n", len(dat))

	if string(dat[0:1]) == "P6" {
		panic("not a PPM P6 image")
	}

	c := 0          // carrage return count
	spaced := false // w x h spaced?
	w := ""         // width
	h := ""         // height
	b := ""         // colour max
	for _, a := range dat[0:14] {
		if a == 10 {
			c++
			continue
		}

		switch c {
		case 1:
			if a == 32 {
				spaced = true
				continue
			}
			if !spaced {
				w += string(a)
			} else {
				h += string(a)
			}
		case 2:
			b += string(a)
		}
	}

	width, err := strconv.Atoi(w)
	if err != nil {
		panic("invalid width " + w)
	}
	height, err := strconv.Atoi(h)
	if err != nil {
		panic("invalid height " + w)
	}
	colours, err := strconv.Atoi(b)
	if err != nil {
		panic("invalid colours " + w)
	}

	fmt.Printf("image w x h: %d x %d\n", width, height)
	fmt.Printf("colour max:  %d\n", colours)

}
