package main

import (
	"fmt"
	"net/http"
	"os"
)

func mains(args []string) error {
	for _, fname := range args {
		fd, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("%s: Open: %s", fname, err)
		}
		var buffer [512]byte
		count, err := fd.Read(buffer[:])
		err2 := fd.Close()
		if err != nil {
			return fmt.Errorf("%s: Read: %w", fname, err)
		}
		if err2 != nil {
			return fmt.Errorf("%s: Close: %w", fname, err2)
		}
		contentType := http.DetectContentType(buffer[:count])
		fmt.Printf("%s\t%s\n", fname, contentType)
	}
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
