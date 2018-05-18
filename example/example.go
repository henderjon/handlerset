package main

import (
	"errors"
	"net/http"

	"github.com/henderjon/knobs"
)

func main() {

	var (
		aMux = http.NewServeMux() // our main mux
		bMux = http.NewServeMux() // paths to require an Authorization header
	)

	aMux.Handle("/foo", knobs.New(
		first(),
		second(),
	))

	aMux.Handle("/bar", knobs.New(
		first(),
		third(),
		second(),
	))

	aMux.Handle("/", knobs.New(
		first(),
		second(),
		third(),
		bMux,
	))

	bMux.Handle("/foo/bar/buzz", fourth())
	bMux.Handle("/foo/bar/bazz", fourth())

	server := &http.Server{
		Handler: aMux,
	}

	server.ListenAndServe()
}

func first() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all is well"))
	})
}

func second() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all is still well"))
	})
}

func third() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("stop asking"))
	})
}

func fourth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		knobs.Cancel(r, errors.New("all is NOT well"))
	})
}
