package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type Header struct {
	HeaderSize      int8
	ProtocolVersion int8
	ProfileVersion  int16
	DataSize        int32
	FormatName      [4]byte
	Crc             uint16
}

func (h Header) String() string {
	return fmt.Sprintf(
		"header size: %d\nprotocolVersion: %d\nprofileVersion: %d\ndataSize: %d\nparsed format: %s\ncrc: %d\n",
		h.HeaderSize,
		h.ProtocolVersion,
		h.ProfileVersion,
		h.DataSize,
		h.FormatName,
		h.Crc,
	)
}

func main() {
	path := "data/data.fit"

	file, fErr := os.Open(path)
	if fErr != nil {
		log.Fatal("Error while opening file", fErr)
	}

	defer file.Close()

	fmt.Printf("%s opened\n", path)

	header := Header{}
	data := readNextBytes(file, 14)
	buffer := bytes.NewBuffer(data)
	bErr := binary.Read(buffer, binary.LittleEndian, &header)
	if bErr != nil {
		log.Fatal("binary.Read failed", bErr)
	}

	fmt.Println(header)
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
