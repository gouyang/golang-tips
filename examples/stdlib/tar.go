package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fw, err := os.Create("files.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	tw := tar.NewWriter(fw)
	defer tw.Close()

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "read me file"},
		{"test.txt", "test file"},
		{"thanks.txt", "thanks all gophers"},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Size: int64(len(file.Body)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}

	r, err := os.Open("files.tar")
	if err != nil {
		log.Fatalln(err)
	}

	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Contents of %s: ", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}
}
