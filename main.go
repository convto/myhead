package main

import (
	"io"
	"os"
	"net/http"
	"flag"
	"bytes"
	"bufio"
	"strings"
	"log"
	"fmt"
)

func readFlag() (int, bool, string) {
	nFlag := flag.Int("n", 10, "number flag")
	remoteFlag := flag.Bool("remote", false, "remote flag")
	flag.Parse()
	path := flag.Arg(0)
	return *nFlag, *remoteFlag, path
}

func getPrintText(reader io.Reader, n int) ([]string, error) {
	var text []string
	var lineCount int
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
	return text, nil
}

func getReader(path string, remote bool) (reader io.Reader, err error) {
	if remote {
		res, err := http.Get(path)
		if err != nil {
			reader = nil
		}
		// *Requestをio.Readerに変換
		// io.Readerならなんでもよかったが、これが適切だと考えた
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		reader = buf
	} else {
		reader, err = os.Open(path)
		if err != nil {
			reader = nil
		}
	}
	return
}

func main() {
	n, remote, path := readFlag()
	if path == "" {
		log.Fatalf("ファイルのパスが入力されていません")
	}
	reader, err := getReader(path, remote)


	fmt.Printf("\n==> %s <==\n", path)
	text, err := getPrintText(reader, n)
	if err != nil {
		log.Fatalf("テキストの読み込みでエラーが発生しました エラー: %v", err)
	}
	fmt.Print(strings.Join(text, "\n"))
	fmt.Printf("\n==> %s end <==\n\n", path)
}