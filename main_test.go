package main

import (
	"testing"
)

var printerBases = []struct {
	isRemote bool
	path string
}{
	{true, "https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt"},
	{false, "./testdata/test.txt"},
	{false, "./testdata/test.png"},
	{false, "./testdata/test.html"},
}

var failPrinterBases = []struct {
	isRemote bool
	path string
}{
	{true, "./testdata/test.txt"},
	{false, "https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt"},
	{false, "not file dir"},
	{false, "./noting.file.ext.example"},
}

func TestArgExists(t *testing.T) {
	err := argsExists(1)
	if err != nil {
		t.Error(err)
	}
}

func TestArgExists2(t *testing.T) {
	err := argsExists(0)
	if err == nil {
		t.Error("prease catch err")
	}
}

func TestGetPrinter(t *testing.T) {
	for _, printerBase := range printerBases {
		_, err := getPrinter(printerBase.isRemote, printerBase.path)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGetPrinter2(t *testing.T) {
	for _, failPrinterBase := range failPrinterBases {
		_, err := getPrinter(failPrinterBase.isRemote, failPrinterBase.path)
		if err == nil {
			t.Error("prease catch err")
		}
	}
}
