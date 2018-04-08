package main

import (
	"flag"
	"os"
	"bufio"
	"log"
	"fmt"
	"strings"
)

func readFlag() (int, string) {
	nFlag := flag.Int("n", 10, "number flag")
	flag.Parse()
	path := flag.Arg(0)
	return *nFlag, path
}

func getPrintText(file *os.File, n int) ([]string, error) {
	var text []string
	var lineCount int
	scanner := bufio.NewScanner(file)
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

func main() {
	n, path := readFlag()
	if path == "" {
		log.Fatalf("ファイルのパスが入力されていません")
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("ファイルの読み込みに失敗しました エラー: %v", err)
	}
	defer file.Close()

	fmt.Printf("\n==> %s <==\n", path)
	text, err := getPrintText(file, n)
	if err != nil {
		log.Fatalf("テキストの読み込みでエラーが発生しました エラー: %v", err)
	}
	fmt.Print(strings.Join(text, "\n"))
	fmt.Printf("\n==> %s end <==\n\n", path)
}