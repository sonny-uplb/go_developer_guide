package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide filename")
		return
	}

	data, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("cant read file:", os.Args[1])
		os.Exit(1)
	}

	// my answer
	b := make([]byte, 1024)

	for {
		n, err := data.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Println(string(b[:n]))
		}
	}

	// from answer of instructor
	// io.Copy(os.Stdout, data)

	data.Close()
}
