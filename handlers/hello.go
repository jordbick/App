package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler that is passed into the ServeHTTP function
type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	// Reading everything from the body and reading into the variable d
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}
