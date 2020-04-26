package messages

import (
	"bytes"
	"encoding/binary"
	"go-fit/internal/io"
	"log"
	"os"
)

type Definition struct {
	Architecture              int8
	GlobalMessageNumber       uint16
	FieldDefinitions          []FieldDefinition
	DeveloperFieldDefinitions []FieldDefinition
}

type FieldDefinition struct {
	Number   uint8
	Size     uint8
	BaseType uint8
}

func NewDefinition(file *os.File) *Definition {
	descriptionBytes := io.ReadNextBytes(file, 5)

	architecture := int8(descriptionBytes[1])

	var globalMessageNumber uint16

	if architecture == 0 {
		globalMessageNumber = binary.LittleEndian.Uint16(descriptionBytes[2:4][:])
	} else {
		globalMessageNumber = binary.BigEndian.Uint16(descriptionBytes[2:4][:])
	}

	numberOfFields := int(descriptionBytes[4])
	fieldsLength := 3 * numberOfFields

	byteFieldsDefinitions := io.ReadNextBytes(file, fieldsLength)
	fieldsBuffer := bytes.NewBuffer(byteFieldsDefinitions)

	var fieldDefinitions []FieldDefinition
	for i := 0; i < fieldsLength; i += 3 {
		fieldDefinition := FieldDefinition{}

		err := binary.Read(fieldsBuffer, binary.BigEndian, &fieldDefinition)
		if err != nil {
			log.Fatal("field definition failed", err)
		}

		fieldDefinitions = append(fieldDefinitions, fieldDefinition)
	}

	numberDevFields := int(io.ReadNextBytes(file, 1)[0])
	devLength := 3 * numberDevFields

	devDefinitions := io.ReadNextBytes(file, devLength)
	devBuffer := bytes.NewBuffer(devDefinitions)
	var devFieldDefinitions []FieldDefinition
	for i := 0; i < numberDevFields; i += 3 {
		fieldDefinition := FieldDefinition{}

		err := binary.Read(devBuffer, binary.BigEndian, &fieldDefinition)
		if err != nil {
			log.Fatal("field definition failed", err)
		}

		devFieldDefinitions = append(devFieldDefinitions, fieldDefinition)
	}

	return &Definition{
		Architecture:              architecture,
		GlobalMessageNumber:       globalMessageNumber,
		FieldDefinitions:          fieldDefinitions,
		DeveloperFieldDefinitions: devFieldDefinitions,
	}
}
