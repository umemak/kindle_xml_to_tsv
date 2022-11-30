package main

import (
	"flag"
	"fmt"
	"os"

	kindlexmltotsv "github.com/umemak/kindle_xml_to_tsv"
)

func main() {
	flag.Parse()
	err := run(flag.Arg(0))
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func run(name string) error {
	res, err := kindlexmltotsv.Convert(name)
	if err != nil {
		return fmt.Errorf("kindlexmltotsv.Convert: %w", err)
	}
	fmt.Println(res)
	return nil
}
