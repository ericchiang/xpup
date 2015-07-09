package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"

	"github.com/ericchiang/xpup/Godeps/_workspace/src/launchpad.net/xmlpath"
)

func fatalf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(2)
}

var (
	omitNewline = true
	inputFile   = ""
)

func init() {
	flag.BoolVar(&omitNewline, "n", false, "")
	flag.StringVar(&inputFile, "f", "", "")
}

func usage() {
	fatalf(`usage: xpup [flags] '[xpath expression]'

flags:

    -f   Read from a given input file rather than from stdin.
    -n   If present xpup will omit the new line at the printed result.

`)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	path, err := xmlpath.Compile(args[0])
	if err != nil {
		fatalf("invalid xpath expression '%s': %v\n", os.Args[1], err)
	}

	decode := func(r io.Reader) (*xmlpath.Node, error) {
		decoder := xml.NewDecoder(r)
		decoder.CharsetReader = func(chset string, input io.Reader) (io.Reader, error) {
			if chset == "" {
				return input, nil
			}
			return charset.NewReader(input, chset)
		}
		return xmlpath.ParseDecoder(decoder)
	}

	var root *xmlpath.Node
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			fatalf("could not open input file: %v\n", err)
		}
		root, err = decode(file)
		file.Close()
		if err != nil {
			fatalf("failed to parse file: %v\n", err)
		}
	} else {
		root, err = decode(os.Stdin)
		if err != nil {
			fatalf("failed to parse stdin: %v\n", err)
		}
	}
	out, ok := path.Bytes(root)
	if !ok {
		fmt.Fprintf(os.Stderr, "[xpup: no items selected]\n")
		return
	}
	if omitNewline {
		os.Stdout.Write(out)
	} else {
		fmt.Printf("%s\n", out)
	}
}
