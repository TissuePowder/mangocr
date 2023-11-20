package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

func init() {
	// Refer to these functions so that goimports is happy before boilerplate is inserted.
	_ = context.Background()
	_ = vision.ImageAnnotatorClient{}
	_ = os.Open
}

func detectText(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	if len(annotations) > 0 {

		text := strings.ReplaceAll(annotations[0].Description, "\n", " ")
		fmt.Fprintf(w, "%s\n", text)

	}

	return nil
}
