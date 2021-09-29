package main

func jheader(w, h int) (out []byte) {
	// 2 byte unsized size/length
	var size []byte

	// jpeg magic byte
	out = []byte{0xff, 0xd8}

	// Define quantization table(s), len
	// B.2.4.1 p39
	dqt := []byte{
		0xff, 0xdb, // marker
		0x00, 0x00, // length
	}
	// luminance
	//   Pq: Quantization table element precision. 0 = 8 bit, 0 = 16 bit
	//   Tq: Quantization table destination identifier. 0-3
	lum := *zigZag(QuantizationTableLuminance75())
	dqt = append(dqt, 0x00) // Pq: 0, Tq: 0
	dqt = append(dqt, []byte(lum[:])...)
	// chrominance
	chm := *zigZag(QuantizationTableLuminance75())
	dqt = append(dqt, 0x01) // Pq: 0, Tq: 1
	dqt = append(dqt, []byte(chm[:])...)
	// update dqt length
	size = i16tob(uint16(len(dqt) - 2))[:] // should be 132
	copy(dqt[2:4], size)

	// height dat
	sh := i16tob(uint16(h))
	// width dat
	sw := i16tob(uint16(w))

	// Baseline DCT, len, precision
	sof := []byte{
		0xff, 0xc0, // SOF0: Baseline DCT. marker
		0x00, 0x00, // Lf: Frame header length
		0x08,         // P: Sample precision. 8 bit
		sh[0], sh[1], // Y: height
		sw[0], sw[1], // X: width
		0x03, // Nf: Number of image components in frame, hardcoded to 3 for colour
	}
	// ref B.2 p36
	// image component Y
	sof = append(sof, []byte{
		0x01, // Ci(8bit): Component identifier
		0x22, // Hi(4bit): Horizontal sampling factor | Vi(4bit):  Vertical sampling factor
		0x00, // Tqi(8bit): Quantization table destination selector
	}...)
	// image component Cb
	sof = append(sof, []byte{0x02, 0x11, 0x01}...)
	// image component Cr
	sof = append(sof, []byte{0x03, 0x11, 0x01}...)
	// update sof length
	size = i16tob(uint16(len(sof)) - 2)[:]
	copy(sof[2:4], size)

	// Define Huffman table(s), len
	dht := []byte{0xff, 0xc4, 0x00, 0x00}

	// Start of scan, len
	sos := []byte{0xff, 0xda, 0x00, 0x00}

	out = append(out, dqt...)
	out = append(out, sof...)
	out = append(out, dht...)
	out = append(out, sos...)

	return
}

// binary - integer conversion functions
// ref https://gist.github.com/chiro-hiro/2674626cebbcb5a676355b7aaac4972d

func i16tob(val uint16) []byte {
	r := make([]byte, 2)
	for i := uint16(0); i < 2; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func btoi16(val []byte) uint16 {
	r := uint16(0)
	for i := uint16(0); i < 2; i++ {
		r |= uint16(val[i]) << (8 * i)
	}
	return r
}
