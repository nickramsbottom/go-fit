package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "data/data.fit"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	defer file.Close()

	fmt.Printf("%s opened\n", path)

	header := readNextBytes(file, 14)

	headerSize := header[0]
	protocolVersion := header[1]
	profileVersion := header[2:4]
	dataSize := header[4:8]
	formatName := header[8:12]
	crc := header[12:14]

	fmt.Printf("header size: %d\n", headerSize)
	fmt.Printf("protocolVersion: %d\n", protocolVersion)
	fmt.Printf("profileVersion: %d\n", binary.LittleEndian.Uint16(profileVersion))
	fmt.Printf("dataSize: %d\n", binary.LittleEndian.Uint32(dataSize))
	fmt.Printf("Parsed format: %s\n", formatName)
	fmt.Printf("crc: %d\n", binary.LittleEndian.Uint16(crc))

	if string(formatName) != ".FIT" {
		log.Fatal("Provided fit file is not in correct format.")
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
