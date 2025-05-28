package filetypes

import (
	"io"
)

type MagicNum struct {
	Offset int
	Algorithm string
	Signature []byte
}

// Maps the magic number of a file (assumed to be its first 8 bytes) to a
// compression algorithm. File types that aren't compressed, or aren't
// recognized as such, are not included.
// 
// Note that tar archives have magic numbers starting at an offset of 257
// bytes.
var compressedMagicNums = []MagicNum{
	{ 257, "archive/tar", []byte{ 0x75, 0x73, 0x74, 0x61, 0x72, 0x00, 0x30, 0x30 } },
	{ 257, "archive/tar", []byte{ 0x75, 0x73, 0x74, 0x61, 0x72, 0x20, 0x20, 0x00 } },
	{ 0, "archive/zip", []byte{ 0x50, 0x4B, 0x03, 0x04 } },
	{ 0, "archive/zip", []byte{ 0x50, 0x4B, 0x05, 0x06 } },
	{ 0, "archive/zip", []byte{ 0x50, 0x4B, 0x07, 0x08 } },
	{ 0, "compress/gzip", []byte{ 0x1F, 0x8B } },
	{ 0, "compress/zlib", []byte{ 0x78, 0x01 } },
	{ 0, "compress/zlib", []byte{ 0x78, 0x01 } },
	{ 0, "compress/zlib", []byte{ 0x78, 0x9C } },
	{ 0, "compress/zlib", []byte{ 0x78, 0xDA } },
	{ 0, "compress/zlib", []byte{ 0x78, 0x20 } },
	{ 0, "compress/zlib", []byte{ 0x78, 0x7D } },
	{ 0, "compress/zlib", []byte{ 0x78, 0xBB } },
	{ 0, "compress/zlib", []byte{ 0x78, 0xF9 } },
	{ 0, "compress/bzip2", []byte{ 0x42, 0x5A, 0x68 } },
}

// Guesses a compression algorithm based on the magic number of the input.
// Returns an empty string if the algorithm cannot be determined.
//
// Note that this function will read the first few hundred bytes of the input.
func CompressionFromMagicNum(input io.Reader) (string, error) {
	header := make([]byte, 270)
	_, err := input.Read(header)
	if err != nil {
		return "", err
	}

	for _, magicNum := range compressedMagicNums {
		start := magicNum.Offset
		end := magicNum.Offset + len(magicNum.Signature)

		actualNum := header[start:end]

		match := true

		for i := range len(actualNum) {
			if magicNum.Signature[i] != actualNum[i] {
				match = false 
				break
			}
		}

		if match {
			return magicNum.Algorithm, nil
		}
	}

	return "", nil
}