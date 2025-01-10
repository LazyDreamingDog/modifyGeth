package cryptoupgrade

import (
	"path/filepath"
)

func gofilePath(fileName string) string {
	return filepath.Join(compressedPath, "src", fileName+".go")
}

func sofilePath(fileName string) string {
	return filepath.Join(compressedPath, "src", fileName+".so")
}
