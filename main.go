package main

import (
	"fmt"
	"io"
	"net/http"
)

// RFC 9112: HTTP/1.1
const fileUrl = "https://www.rfc-editor.org/rfc/rfc9112.txt"

func main() {
	fmt.Printf("downloading file from: %s\n", fileUrl)
	fileContent := downloadFile(fileUrl)
	fmt.Printf("content downloaded! (%d byte)\n", len(fileContent))
	length := len(fileContent)
	fmt.Printf("length: %d\n", length)

	_, allChars := createCharAnalysisList(fileContent)
	fmt.Printf("total chars: %d\n", len(allChars))

	entropy := getEntropy(allChars)
	fmt.Printf("entropy: %f\n", entropy)

	encoding := createEncodingHuffman(allChars)
	codes := encoding.coding()
	totalBit := 0
	for _, c := range allChars {
		var d string
		if c.char == '\n' {
			d = "\\n"
		} else {
			d = string(c.char)
		}
		code := codes[c.char]
		bit := len(codes[c.char])
		totalBit += c.count * bit
		fmt.Printf("%s: %d(%.2f%%) x %s(%d bit) = %d bit\n", d, c.count, c.probability*100, code, bit, c.count*bit)
	}
	fmt.Printf("total %d bit (%d byte)\n", totalBit, totalBit/8)
}

func downloadFile(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
