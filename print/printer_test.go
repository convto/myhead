package print_test

import (
	"testing"
	"strings"

	"github.com/srttk/myhead/print"
)

type PrinterBase struct {
	path string
	text string
	n    int
}

var printerBases = []PrinterBase {
	{"%a", "[%a]", 1}, {"%-a", "[%-a]", 100}, {"%+a", "[%+a]", 10000},
	{"%#a", "[%#a]", 50}, {"% a", "[% a]", 300},
}

func testGetPrintText(t *testing.T, printerBase PrinterBase) {
	t.Helper()
	reader := strings.NewReader(printerBase.text)
	_, err := print.GetPrintText(reader, printerBase.n, printerBase.path)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPrintText(t *testing.T) {
	for _, readerBase := range printerBases {
		testGetPrintText(t, readerBase)
	}
}

func testPrintHepler(t *testing.T, printerBase PrinterBase, printer print.Printer) {
	t.Helper()
	err := printer.Print(printerBase.n, printerBase.path)
	if err != nil {
		t.Error(err)
	}
}
