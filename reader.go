package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Reader struct {
	csvReader *csv.Reader
}

func StdinReader() *Reader {
	return &Reader{
		csvReader: csv.NewReader(os.Stdin),
	}
}

func FileReader(filePath string) (*Reader, error) {

	f, err := os.Open(filePath)

	var r *Reader

	if err == nil {
		r = &Reader{
			csvReader: csv.NewReader(f),
		}
	}

	return r, err
}

func (r *Reader) ReadLine() (string, string, error) {

	segments, err := r.csvReader.Read()

	if err == io.EOF {
		return "", "", err
	}

	if len(segments) < 2 {
		return "", "", fmt.Errorf("Line does not contain two columns, skipping")
	}

	return segments[0], segments[1], nil
}
