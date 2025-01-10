package cryptoupgrade

import (
	"os"
	"path/filepath"
)

func gofilePath(fileName string) string {
	return filepath.Join(compressedPath, "src", fileName+".go")
}

func sofilePath(fileName string) string {
	return filepath.Join(compressedPath, "so", fileName+".so")
}

func directoryInit() {
	os.MkdirAll(filepath.Join(compressedPath, "src"), os.ModePerm)
	os.MkdirAll(filepath.Join(compressedPath, "so"), os.ModePerm)
}
