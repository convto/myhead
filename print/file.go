package print

import (
	"io"
	"os"
	"bytes"
	"fmt"
	"strings"
	"path/filepath"
)

type File struct {
	io.Reader
}

func NewFile(path string) (File, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	file, err := os.Open(fullPath)
	defer file.Close()
	if err != nil {
		return File{}, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		return File{}, err
	}
	return File{buf}, nil
}

func (f File) Print(n int, path string) error {
	fileName := filepath.Base(path)
	strSlice, err := getPrintText(f, n, fileName)
	if err != nil {
		return err
	}
	fmt.Print(strings.Join(strSlice, "\n"))
	return nil
}

