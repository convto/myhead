package print_test

import (
	"testing"

	"github.com/srttk/myhead/print"
	"strings"
)

var paths = []string{"../testdata/test.txt", "../testdata/test.png", "../testdata/test.html"}
var failPaths = []string{"not dir character", "./noting.file.ext.example", }

func TestNewFile(t *testing.T) {
	var err error
	for _, path := range paths {
		_, err = print.NewFile(path)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNewFile2(t *testing.T) {
	var err error
	for _, failPath := range failPaths {
		_, err = print.NewFile(failPath)
		if err == nil {
			t.Error("prease catch err")
		}
	}
}

func TestFile_Print(t *testing.T) {
	for _, fileBase := range printerBases {
		reader := strings.NewReader(fileBase.text)
		file := print.File{ Reader: reader }
		testPrintHepler(t, fileBase, file)
	}

}