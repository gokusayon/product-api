package handlers

import (
	"log"
	"net/http"
)

type Ping struct {
	log *log.Logger
}

func NewPing(log *log.Logger) *Ping {
	return &Ping{log}
}

func (h *Ping) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.log.Println("Running Ping Handler")
	rw.Write([]byte("Ping Successfull!!"))
}
