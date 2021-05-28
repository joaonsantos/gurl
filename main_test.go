package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, `<h1>hello</h1>`)
	})
}

func TestGurl(t *testing.T) {
	testServer := httptest.NewServer(handler())
	defer testServer.Close()

	out := httptest.NewRecorder()
	args := Args{
		Headers: false,
		URL:     testServer.URL,
	}

	err := Fetch(out, args)
	if err != nil {
		t.Errorf("did not expect an error, got: %v", err)
	}

	if size := len(out.Body.Bytes()); size == 0 {
		t.Errorf("did not expect content to be empty, got size %v", size)
	}
}
