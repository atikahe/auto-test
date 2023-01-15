package testfile

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Generate(filePath string) (string, error) {
	fileName := filepath.Base(filePath)
	ext := filepath.Ext(fileName)
	base := strings.TrimSuffix(fileName, ext)

	dir := filepath.Dir(filePath) + "/"

	switch ext {
	case ".java":
		return dir + base + "Test" + ext, nil
	case ".js":
		return dir + base + ".test" + ext, nil
	case ".py":
		return dir + "test_" + base + ext, nil
	case ".cs":
		return dir + base + "Tests" + ext, nil
	case ".go":
		return dir + base + "_test" + ext, nil
	default:
		return "", fmt.Errorf("unknown file extension %s", ext)
	}
}
