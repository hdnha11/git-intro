package main

import (
	"compress/zlib"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	filePath := flag.String("f", "", "zlib format compressed file path")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("cannot open input file: %v\n", err)
	}
	defer file.Close()

	r, err := zlib.NewReader(file)
	if err != nil {
		log.Fatalf("cannot create zlib reader: %v\n", err)
	}
	defer r.Close()

	io.Copy(os.Stdout, r)
}
