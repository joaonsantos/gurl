package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCurlErrors(t *testing.T) {

	tt := []struct {
		name string
		args []string
		err  string
	}{
		{
			"no args",
			[]string{"main"},
			"no arguments",
		},
		{
			"no protocol",
			[]string{"main", "google.com"},
			`Get "google.com": unsupported protocol scheme ""`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := httptest.NewRecorder()

			err := Fetch(out, tc.args)
			if err.Error() != tc.err {
				t.Errorf("expected to get %v, got %v", tc.err, err)
			}
		})
	}

}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, `<h1>hello</h1>`)
	})
}

func TestCurlOk(t *testing.T) {
	testServer := httptest.NewServer(handler())
	defer testServer.Close()

	out := httptest.NewRecorder()
	args := []string{"main", testServer.URL}

	err := Fetch(out, args)
	if err != nil {
		t.Errorf("did not expect an error, got: %v", err)
	}

	if size := len(out.Body.Bytes()); size == 0 {
		t.Errorf("did not expect content to be empty, got size %v", size)
	}
}
