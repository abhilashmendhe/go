package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello, World")
	d, err := ioutil.ReadAll(r.Body) // will return some data and err. from user
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s\n", d)
}