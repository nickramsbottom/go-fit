package main

import (
	"go-fit/internal/headers"
	"go-fit/internal/io"

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

	recordHeaderByte := io.ReadNextBytes(file, 1)[0]
	fmt.Printf("%08b\n", recordHeaderByte)

	normalHeader := GetBit(recordHeaderByte, 7)
	messageType := GetBit(recordHeaderByte, 6)
	specific := GetBit(recordHeaderByte, 5)
	reserved := GetBit(recordHeaderByte, 4)
	localMessageType := LocalMessageType(recordHeaderByte)

	fmt.Println(normalHeader)
	fmt.Println(messageType)
	fmt.Println(specific)
	fmt.Println(reserved)
	fmt.Println(localMessageType)
}

func LocalMessageType(b byte) int {
	return int(b & 0xF)
}

func GetBit(b byte, bitNumber int) bool {
	return (b & (1 << bitNumber)) != 0
}
