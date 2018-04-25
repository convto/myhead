package main

import (
	"os"
	"flag"
	"fmt"

	"github.com/srttk/myhead/print"
)

var (
	n        int
	isRemote bool
	path     string

	printer  print.Printer
	err      error
)

func init() {
	flag.IntVar(&n, "n", 10, "|optional| number of rows to read. default value 10")
	flag.BoolVar(&isRemote, "remote", false, "|optional| whether to read external files using http. If true it will be treated as http resource")
	flag.BoolVar(&isRemote, "r", false, "|optional| 'remote' shorthand")
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Print("required 'path' params\nPlease start with -h and check the required items")
		os.Exit(2)
	}
	path = flag.Arg(0)

	if isRemote {
		printer, err = print.NewResponse(path)
	} else {
		printer, err = print.NewFile(path)
	}
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