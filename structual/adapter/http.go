package adapter

import (
	"fmt"
	"net/http"
)

type MyServer struct{
	Msg string
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.Msg)
}
