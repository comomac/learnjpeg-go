package main

type TypeQuantizationTable [64]int

// ref CCITTRec. T.81 (1992 E) Table K.1 p.143

// Quantization Table Luminance at 75% quality
func QuantizationTableLuminance75() *TypeQuantizationTable {
	return zigZag(&TypeQuantizationTable{
		16, 11, 10, 16, 24, 40, 51, 61, // 0-7
		12, 12, 14, 19, 26, 58, 60, 55, // 8-15
		14, 13, 16, 24, 40, 57, 69, 56, // 16-23
		14, 17, 22, 29, 51, 87, 80, 62, // 24-31
		18, 22, 37, 56, 68, 109, 103, 77, // 32-39
		24, 35, 55, 64, 81, 104, 113, 92, // 40-47
		49, 64, 78, 87, 103, 121, 120, 101, // 48-55
		72, 92, 95, 98, 112, 100, 103, 99, // 56-63
	})
}

// ref CCITTRec. T.81 (1992 E) Table K.2 p.143

// Quantization Table Chrominance at 75% quality
func QuantizationTableChrominance() *TypeQuantizationTable {
	return zigZag(&TypeQuantizationTable{
		17, 18, 24, 47, 99, 99, 99, 99,
		18, 21, 26, 66, 99, 99, 99, 99,
		24, 26, 56, 99, 99, 99, 99, 99,
		47, 66, 99, 99, 99, 99, 99, 99,
		99, 99, 99, 99, 99, 99, 99, 99,
		99, 99, 99, 99, 99, 99, 99, 99,
		99, 99, 99, 99, 99, 99, 99, 99,
		99, 99, 99, 99, 99, 99, 99, 99,
	})
}

// zig zag order
func zigZag(in *TypeQuantizationTable) (out *TypeQuantizationTable) {
	out = &TypeQuantizationTable{
		in[0], in[1], in[8], in[16], in[9], in[2], in[3], in[10],
		in[17], in[24], in[32], in[25], in[18], in[11], in[4], in[5],
		in[12], in[19], in[26], in[33], in[40], in[48], in[41], in[34],
		in[27], in[20], in[13], in[6], in[7], in[14], in[21], in[28],
		in[35], in[42], in[49], in[56], in[57], in[50], in[43], in[36],
		in[29], in[22], in[15], in[23], in[30], in[37], in[44], in[51],
		in[58], in[59], in[52], in[45], in[38], in[31], in[39], in[46],
		in[53], in[60], in[61], in[54], in[47], in[55], in[62], in[63],
	}
	return out
}
