package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.log.Println("Running Hello Handler")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Hello %s", b)

}
