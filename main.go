package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Args struct {
	Headers bool
	URL     string
}

func addProto(URL string) string {
	protoURL := URL

	match, _ := regexp.MatchString(".*://.*", URL)
	if !match {
		fullURLSlice := append([]string{"http://"}, URL)
		protoURL = strings.Join(fullURLSlice, "")
	}

	return protoURL
}

func printHeaders(out io.Writer, r *http.Response) {
	fmt.Fprintf(out, "%s %s\n", r.Proto, r.Status)

	for k, v := range r.Header {
		fmt.Fprintf(out, "%s: %s\n", k, v[0])
	}
}

func Fetch(out io.Writer, args Args) error {
	args.URL = addProto(args.URL)
	res, err := http.Get(args.URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if args.Headers {
		printHeaders(os.Stdout, res)
	}
	if _, err := io.Copy(out, res.Body); err != nil {
		return err
	}
	return nil
}

func usage() {
	fmt.Print("Usage: gurl [options...] <url>\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage // override usage func
	incHeaders := flag.Bool("i", false, "include protocol response headers")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}
	args := Args{*incHeaders, flag.Arg(0)}

	if err := Fetch(os.Stdout, args); err != nil {
		fmt.Fprintf(os.Stderr, "gurl: could not complete request: %v\n", err)
		os.Exit(2)
	}
}
