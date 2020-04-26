package headers

import (
	"go-fit/internal/io"
	"os"
)

type RecordHeader struct {
	NormalHeader     bool
	MessageType      bool
	MessageSpecific  bool
	LocalMessageType int
}

func NewRecordHeader(file *os.File) *RecordHeader {
	headerByte := io.ReadNextBytes(file, 1)[0]

	return &RecordHeader{
		NormalHeader:     getBit(headerByte, 7),
		MessageType:      getBit(headerByte, 6),
		MessageSpecific:  getBit(headerByte, 5),
		LocalMessageType: int(headerByte & 0xF),
	}
}

func getBit(b byte, bitNumber int) bool {
	return (b & (1 << bitNumber)) != 0
}
