package main

func jheader(w, h uint32) (out []byte) {

	// jpeg magic byte
	out = []byte{0xff, 0xd8}

	// Define quantization table(s), len
	dqt := []byte{0xff, 0xdb, 0x00, 0x84} // 0x84 == 132, pre-calculated
	// luminance
	lum := *zigZag(QuantizationTableLuminance75())
	dqt = append(dqt, 0x00) // Pq: 0, Tq: 0
	dqt = append(dqt, []byte(lum[:])...)
	// chrominance
	chm := *zigZag(QuantizationTableLuminance75())
	dqt = append(dqt, 0x01) // Pq: 0, Tq: 1
	dqt = append(dqt, []byte(chm[:])...)

	// Baseline DCT, len
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
