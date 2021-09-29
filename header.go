package main

func jheader(w, h uint16) (out []byte) {
	// 2 byte unsized size/len
	var size []byte

	// jpeg magic byte
	out = []byte{0xff, 0xd8}

	// Define quantization table(s), len
	dqt := []byte{0xff, 0xdb, 0x00, 0x0}
	// luminance
	lum := *zigZag(QuantizationTableLuminance75())
	dqt = append(dqt, 0x00) // Pq: 0, Tq: 0
	dqt = append(dqt, []byte(lum[:])...)
	// chrominance
	chm := *zigZag(QuantizationTableLuminance75())
	dqt = append(dqt, 0x01) // Pq: 0, Tq: 1
	dqt = append(dqt, []byte(chm[:])...)
	size = i16tob(uint16(len(dqt) - 2))[:] // should be 132
	// update dqt size
	copy(dqt[2:4], size)

	// Baseline DCT, len, precision
	sof := []byte{0xff, 0xc0, 0x00, 0x00}

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
