package main

import (
	"testing"
	"strings"
	"io"
)

func TestGetReader(t *testing.T) {
	var err error

	// local
	_, err = getReader("./test.txt", false)
	if err != nil {
		t.Errorf("ローカルファイルの読み込みに失敗しました エラー: %v", err)
	}

	// remote
	_, err = getReader("https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt", true)
	if err != nil {
		t.Errorf("リモートファイルの読み込みに失敗しました エラー: %v", err)
	}
}

func TestGetPrintText(t *testing.T) {
	var (
		reader io.Reader
		text   []string
		str    string
		err    error
	)
	// local
	reader, _ = getReader("./test.txt", false)
	text, err = getPrintText(reader, 2)
	if err != nil {
		t.Errorf("テキストの読み込みでエラーが発生しました エラー: %v", err)
	}
	// getPrintTextの結果を評価するために[]stringをstringに変換
	str = strings.Join(text, "\n")
	if str != "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit," {
		t.Errorf("期待される出力と差があります")
	}

	// remote
	reader, _ = getReader("https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt", true)
	text, err = getPrintText(reader, 2)
	if err != nil {
		t.Errorf("テキストの読み込みでエラーが発生しました エラー: %v", err)
	}
	// getPrintTextの結果を評価するために[]stringをstringに変換
	str = strings.Join(text, "\n")
	if str != "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit," {
		t.Errorf("期待される出力と差があります")
	}
}
