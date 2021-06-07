package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"strconv"
)

// PPMOpenFile open PPM file and decode it
func PPMOpenFile(fileName string) (numberOfColours int, img *image.RGBA, err error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return 0, nil, err
	}

	// fmt.Printf("filesize: %d bytes\n", len(dat))

	return PPMDecode(dat)
}

// PPMOpenFile open PPM bytes and decode it
func PPMDecode(dat []byte) (numberOfColours int, img *image.RGBA, err error) {
	if string(dat[0:1]) == "P6" {
		return 0, nil, fmt.Errorf("not a PPM P6 image")
	}

	c := 0          // carrage return count
	spaced := false // w x h spaced?
	w := ""         // width
	h := ""         // height
	b := ""         // colour max
	q := 0          // position of first image byte
outer:
	for i, a := range dat[0:14] {
		q = i
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
		case 3:
			break outer
		}
	}

	width, err := strconv.Atoi(w)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid width " + w)
	}
	height, err := strconv.Atoi(h)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid height " + w)
	}
	numberOfColours, err = strconv.Atoi(b)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid colours " + w)
	}

	img = image.NewRGBA(image.Rect(0, 0, width, height))
	x := 0 // pixel x position
	y := 0 // pixel y position
	idat := dat[q:]
	px := []byte{} // 3 bytes makes a pixel
	for _, b := range idat {
		x++
		if b == 10 {
			x = 0
			y++
			continue
		}
		px = append(px, b)
		if len(px) == 3 {
			img.SetRGBA(x, y, color.RGBA{
				R: px[0],
				G: px[1],
				B: px[2],
			})
			px = []byte{}
		}
	}

	return numberOfColours, img, nil
}
