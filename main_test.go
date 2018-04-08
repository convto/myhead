package main

import (
	"testing"
	"os"
	"strings"
)

func TestGetPrintText(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		t.Errorf("ファイルの読み込みに失敗しました エラー: %v", err)
	}
	defer file.Close()

	text, err := getPrintText(file, 2)
	if err != nil {
		t.Errorf("テキストの読み込みでエラーが発生しました エラー: %v", err)
	}
	str := strings.Join(text, "\n")
	if str != "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit," {
		t.Errorf("期待される出力と差があります")
	}
}
