package main

// Pixel holds smallest unit of image
type Pixel [3]byte

func rgbToYCbCr(pixelRGB Pixel) (pixelYCbCr Pixel) {
	var r, g, b *byte

	r = &pixelRGB[0]
	g = &pixelRGB[1]
	b = &pixelRGB[2]

	// Reference: jfif3.pdf - Conversion to and from RGB
	pixelYCbCr[0] = byte(0.299*float32(*r)) + byte(0.587*float32(*g)) + byte(0.114*float32(*b))
	pixelYCbCr[1] = byte(-0.1687*float32(*r)) - byte(0.3313*float32(*g)) + byte(0.5*float32(*b)) + 128
	pixelYCbCr[2] = byte(0.5*float32(*r)) - byte(0.4187*float32(*g)) - byte(0.0813*float32(*b)) + 128

	return
}

func YCbCrToRGB(pixelYCbCr Pixel) (pixelRGB Pixel) {
	var y, cb, cr *byte

	y = &pixelYCbCr[0]
	cb = &pixelYCbCr[1]
	cr = &pixelYCbCr[2]

	// Reference: jfif3.pdf - Conversion to and from RGB
	pixelRGB[0] = *y + byte(1.402*float32(*cr-128))
	pixelRGB[1] = *y - byte(0.34414*(float32(*cb)-128)) - byte(0.71414*float32(*cr-128))
	pixelRGB[2] = *y + byte(1.772*(float32(*cb)-128))

	return
}
