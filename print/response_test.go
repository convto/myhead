package print_test

import (
	"testing"

	"github.com/srttk/myhead/print"
	"strings"
)

func TestNewResponse(t *testing.T) {
	_, err := print.NewResponse("https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt")
	if err != nil {
		t.Error(err)
	}
}

func TestResponse_Print(t *testing.T) {
	for _, responseBase := range printerBases {
		reader := strings.NewReader(responseBase.text)
		response := print.Response{ Reader: reader }
		testPrintHepler(t, responseBase, response)
	}
}