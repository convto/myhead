package main

import (
	"os"
	"flag"
	"fmt"
	"errors"

	"github.com/srttk/myhead/print"
)

var (
	n        int
	isRemote bool
	path     string

	printer  print.Printer
	err      error
)

func argsExists(argsLen int) error {
	if argsLen == 0 {
		return errors.New("required 'path' params\nPlease start with -h and check the required items")
	}
	return nil
}

func getPrinter(isRemote bool, path string) (print.Printer, error) {
	if isRemote {
		printer, err = print.NewResponse(path)
	} else {
		printer, err = print.NewFile(path)
	}
	return printer, err
}

func main() {
	flag.IntVar(&n, "n", 10, "|optional| number of rows to read. default value 10")
	flag.BoolVar(&isRemote, "remote", false, "|optional| whether to read external files using http. If true it will be treated as http resource")
	flag.BoolVar(&isRemote, "r", false, "|optional| 'remote' shorthand")
	flag.Parse()
	if err := argsExists(len(flag.Args())); err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	path = flag.Arg(0)

	printer, err := getPrinter(isRemote, path)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	err = printer.Print(n, path)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
}