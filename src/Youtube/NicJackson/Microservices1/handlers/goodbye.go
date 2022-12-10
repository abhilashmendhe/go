package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (gh *Goodbye) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	rw.Write([]byte("Byeeee..\n"))
}
