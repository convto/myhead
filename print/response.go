package print

import (
	"io"
	"bytes"
	"fmt"
	"strings"
	"net/http"
)

type Response struct {
	io.Reader
}

func NewResponse(url string) (Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return Response{}, err
	}
	return Response{buf}, nil
}

func (r Response) Print(n int, path string) error {
	strSlice, err := GetPrintText(r.Reader, n, path)
	if err != nil {
		return err
	}
	fmt.Print(strings.Join(strSlice, "\n"))
	return nil
}
