package main

import (
	"flag"
	"os"
	"bufio"
	"log"
	"fmt"
)

func readFlag() (n int, path string) {
	// パースしてからでないと使えない関数、そうでない関数がflagに存在する可能性がある
	// os.Argsの方がいいかも
	n = *flag.Int("n", 10, "number flag")
	flag.Parse()
	path = flag.Arg(0)
	return
}

func printFile(path string, file *os.File, n int) {
	fmt.Printf("\n==> %s <==\n", path)
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCount += 1
		if lineCount >= n {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("==> %s end <==\n\n", path)
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
	printFile(path, file, n)
}