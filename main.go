package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch(out io.Writer, args []string) error {
	if len(args) < 2 {
		return errors.New("no arguments")
	}
	url := args[1]

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	io.Copy(out, res.Body)
	return nil
}

func main() {
	args := os.Args
	err := Fetch(os.Stdout, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not fetch url: %v\n", err)
		os.Exit(1)
	}
}
