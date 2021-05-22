# Learn JPEG in Go

Learn how JPEG works by reinventing the library. Focus will be on bare minimum amount of code and the workings of code are understandable.

## Limitation (to speed up dev, learn and reduce complexity)

* Encode only
* Portable Pixmap Format P6 input file format only
* Quality set to 75%, no other quality
* No Chroma subsampling
* Width and Height are divisible by 16
* No optimisation
* Use precomputed or provided numbers

## Order

* Convert RGB to YCbCr
* ~~Chroma Subsampling 4:2:0~~
* Pixels block to 8x8
* Discrete Cosine Transform (DCT II)
* Quantization
* Encode
    * Zigzag
    * Huffman
    * Run-length encoding
* Store Header + Data

## References 

* [Wikipedia JPEG](https://en.wikipedia.org/wiki/JPEG)
* [Wikipedia Netpbm](https://en.wikipedia.org/wiki/Netpbm)
* [JPEG File Interchange Format](https://en.wikipedia.org/wiki/JPEG_File_Interchange_Format)
* [JPEG Header Information](http://www.geocities.ws/crestwoodsdd/JPEG.htm)
* [YCbCr](https://en.wikipedia.org/wiki/YCbCr#JPEG_conversion)
* [Chroma subsampling](https://en.wikipedia.org/wiki/Chroma_subsampling#Sampling_systems_and_ratios)
* [Huffman Coding](https://en.wikipedia.org/wiki/Huffman_coding)
* [Run-length Encoding](https://en.wikipedia.org/wiki/Run-length_encoding)
* [JPEG File Layout and Format](http://vip.sugovica.hu/Sardi/kepnezo/JPEG%20File%20Layout%20and%20Format.htm)
* [Go JPEG library](https://golang.org/pkg/image/jpeg)
* [Colourspaces (JPEG Pt0)- Computerphile](https://www.youtube.com/watch?v=LFXN9PiOGtY)
* [JPEG 'files' & Colour (JPEG Pt1)- Computerphile](https://www.youtube.com/watch?v=n_uNPbdenRs)
* [JPEG DCT, Discrete Cosine Transform (JPEG Pt2)- Computerphile](https://www.youtube.com/watch?v=Q2aEzeMDHMA)
* [ITU-T T.81](https://www.w3.org/Graphics/JPEG/itu-t81.pdf)
* [JPEG File Interchange Format](https://www.w3.org/Graphics/JPEG/jfif3.pdf)
* [JPEG Standard](https://web.stanford.edu/class/ee398a/handouts/lectures/08-JPEG.pdf)
* [Chapter 9 Image Compression Standards - Rutgers CS](https://www.cs.rutgers.edu/~elgammal/classes/cs334/slide9_short.pdf)
