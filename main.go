package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "keyfile.json")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <path-to-image-or-directory>\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "Pass a path to a local file.\n")
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)
	fi, err := os.Stat(path)

	images := []string{}

	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	if fi.IsDir() {
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {
			images = append(images, filepath.Join(path, entry.Name()))
		}
	} else {
		images = append(images, path)
	}

	// fmt.Println(images)

	for _, image := range images {

		txt := strings.TrimSuffix(filepath.Base(image), filepath.Ext(image)) + ".txt"
		txtPath := filepath.Join("ocr", txt)
		txtfile, err := os.OpenFile(txtPath, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			log.Fatal(err)
		}

		if err := detectTextLocal(txtfile, image); err != nil {
			log.Fatal(err)
		}
		if err := os.Remove(image); err != nil {
			log.Fatal(err)
		}

		fmt.Println(txtPath)
	}

}
