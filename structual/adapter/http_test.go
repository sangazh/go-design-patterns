package adapter

import (
	"log"
	"net/http"
	"testing"
)

func TestHTTP(t *testing.T) {
	server := &MyServer{Msg: "Hello, World"}

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
