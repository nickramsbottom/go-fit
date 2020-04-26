package main

import (
	"go-fit/internal/headers"

	"fmt"
	"log"
	"os"
)

type ContentHeader struct {
	Reserved            int8
	Architecture        int8
	GlobalMessageNumber int16
	NumberOfFields      int8
}

func main() {
	path := "data/data.fit"

	file, fErr := os.Open(path)
	if fErr != nil {
		log.Fatal("Error while opening file", fErr)
	}

	defer file.Close()

	fmt.Printf("%s opened\n", path)

	fileHeader := headers.NewFileHeader(file)
	fmt.Println(fileHeader)

	recordHeader := headers.NewRecordHeader(file)
	fmt.Printf("%+v\n", recordHeader)
	// fmt.Printf("%08b\n", recordHeaderByte)
}

func LocalMessageType(b byte) int {
	return int(b & 0xF)
}

func GetBit(b byte, bitNumber int) bool {
	return (b & (1 << bitNumber)) != 0
}
