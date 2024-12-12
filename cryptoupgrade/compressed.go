package cryptoupgrade

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
)

// * Serialize .go file to string
func compressFileToString(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// use gzip to compress
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(fileData)
	if err != nil {
		return "", err
	}
	writer.Close()
	// Use base64 encoding
	compressedData := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return compressedData, nil
}

// * Deserialize string to .go file
func decompressStringToFile(compressedString string, outputPath string) error {
	// Decode base64
	compressedData, err := base64.StdEncoding.DecodeString(compressedString)
	if err != nil {
		return err
	}
	// Use gzip to zip
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return err
	}
	defer reader.Close()
	decodedData, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	// Output
	err = os.WriteFile(outputPath, decodedData, 0644)
	if err != nil {
		return err
	}
	return nil
}