# JPEG format minimum

  *B.2.1  p34

Code Assignment
      Symbol
            Description

FFD8  SOI   Start Of Image

  * B.2.4.1  p39
FFDB  DQT   Define quantization table(s)
xxxx        Size of quantization table(s) (whole DQT minus Code Assignment)
xxxx[n]...  Content of quantization table[n] (65 bytes each)

  * B2.2  p35
FFC0  SOF0  Baseline DCT
xxxx        Size of DCT (whole SOF0 minus Code Assignment)
xxxx...     Content of DCT

  * B.2.4.2  p40
FFC4  DHT   Define Huffman table(s)
xxxx        Size of Huffman table (whole DHT minus Code Assignment)
xxxx[n]...  Content of Huffman table[n]

  * B2.3  p37
FFDA  SOS   Start of scan
xxxx        Size of Start of scan (whole SOS minus Code Assignment and Image Data)
xxxx[n]...  Content of Scan Header[n] (2 bytes each)
xxxx...     Image Data

FFD9  EOI   End of Image
