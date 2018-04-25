package print

import (
	"io"
	"bufio"
)


type Printer interface {
	Print(int, string) error
}

func GetPrintText(reader io.Reader, n int, path string) ([]string, error) {
	var text []string
	var lineCount int
	text = append(text, "\n==>"+path+"<==")
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text = append(text, scanner.Text())
		lineCount += 1
		if lineCount >= n {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	text = append(text, "==>"+path+" end"+"<==\n\n")
	return text, nil
}
