package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "keyfile.json")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <path-to-image>\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "Pass either a path to a local file.\n")
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)

	sample := struct {
		name  string
		local func(io.Writer, string) error
	}{"detectText", detectText}

	err := sample.local(os.Stdout, path)

	if err != nil {
		fmt.Println("Error:", err)
	}

}
