package tmps

//
//import (
//	"bytes"
//	"encoding/base64"
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"path/filepath"
//
//	"github.com/klauspost/compress/gzip"
//)
//
//// FileData represents the encoded file data
//type FileData struct {
//	Type     string                 `json:"type"`
//	Metadata map[string]interface{} `json:"metadata"`
//	Data     string                 `json:"data"`
//}
//
//// FileToData 将文件转换为 FileData
//func FileToData(filePath string) (*FileData, error) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		return nil, fmt.Errorf("无法打开文件：%w", err)
//	}
//	defer file.Close()
//
//	fileBytes, err := io.ReadAll(file)
//	if err != nil {
//		return nil, fmt.Errorf("无法读取文件内容：%w", err)
//	}
//
//	mimeType := http.DetectContentType(fileBytes)
//	fileExt := filepath.Ext(filePath)
//
//	metadata := map[string]interface{}{
//		"filename":  filepath.Base(filePath),
//		"filesize":  len(fileBytes),
//		"extension": fileExt,
//	}
//
//	encodedData := base64.StdEncoding.EncodeToString(fileBytes)
//
//	fileData := &FileData{
//		Type:     mimeType,
//		Metadata: metadata,
//		Data:     encodedData,
//	}
//
//	return fileData, nil
//}
//
//func DataToFile(filePath string, fileData *FileData) error {
//	// 解码 Base64 编码的数据
//	decodedData, err := base64.StdEncoding.DecodeString(fileData.Data)
//	if err != nil {
//		return fmt.Errorf("无法解码 Base64 数据：%w", err)
//	}
//
//	// 创建文件
//	file, err := os.Create(filePath)
//	if err != nil {
//		return fmt.Errorf("无法创建文件：%w", err)
//	}
//	defer file.Close()
//
//	_, err = file.Write(decodedData)
//	if err != nil {
//		return fmt.Errorf("无法写入文件内容：%w", err)
//	}
//
//	return nil
//}
//
//func compressData(data []byte) ([]byte, error) {
//	var buf []byte
//	gzipWriter, err := gzip.NewWriterLevel(io.Writer(&buf), gzip.BestCompression)
//	if err != nil {
//		return nil, err
//	}
//	_, err = gzipWriter.Write(data)
//	if err != nil {
//		return nil, err
//	}
//	err = gzipWriter.Close()
//	if err != nil {
//		return nil, err
//	}
//
//	return buf, nil
//}
//
//func decompressData(data []byte) ([]byte, error) {
//	buf := new(bytes.Buffer)
//	reader, err := gzip.NewReader(bytes.NewReader(data))
//	if err != nil {
//		return nil, err
//	}
//	defer reader.Close()
//
//	_, err = io.Copy(buf, reader)
//	if err != nil {
//		return nil, err
//	}
//
//	return buf.Bytes(), nil
//}
