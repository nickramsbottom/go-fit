package main

import (
	"go-fit/internal/headers"
	"go-fit/internal/messages"

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

	definitionMessage := messages.NewDefinition(file)
	fmt.Printf("%+v\n", definitionMessage)
}
