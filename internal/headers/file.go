package headers

import (
	"go-fit/internal/io"

	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type FileHeader struct {
	HeaderSize      int8
	ProtocolVersion int8
	ProfileVersion  int16
	DataSize        int32
	FormatName      [4]byte
	Crc             uint16
}

func NewFileHeader(f *os.File) *FileHeader {
	header := FileHeader{}
	data := io.ReadNextBytes(f, 14)
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.LittleEndian, &header)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	return &header
}

func (h FileHeader) String() string {
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
