package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple Handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new Hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}

// ServeHTTP implements the go http.Handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
