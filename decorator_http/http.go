package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator")
}

func Start() {
	http.Handle("/", &LoggerServer{
		Handler:   new(MyServer),
		LogWriter: os.Stdout,
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI: %s \n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprintln(s.LogWriter, "-------")
	s.Handler.ServeHTTP(w, r)
}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok {
		fmt.Fprintln(w, "error retrive data from Basic auth")
		return
	}

	if user == s.User && pass == s.Password {
		s.Handler.ServeHTTP(w, r)
	} else {
		fmt.Fprintln(w, "incorrect user or password")
	}
}

func main() {
	fmt.Println("enter the type number of server you want to launch from the following:")
	fmt.Println("1: Plain server")
	fmt.Println("2: Server with logging")
	fmt.Println("3: Server with logging and authentication")

	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)
	var mySuperServer http.Handler

	switch selection {
	case 1:
		mySuperServer = new(MyServer)
	case 2:
		mySuperServer = &LoggerServer{
			Handler:   new(MyServer),
			LogWriter: os.Stdout,
		}
	case 3:
		var user, password string
		fmt.Println("Enter user and password separated by a space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)
		mySuperServer = &LoggerServer{
			Handler: &BasicAuthMiddleware{
				Handler:  new(MyServer),
				User:     user,
				Password: password,
			},
			LogWriter: os.Stdout,
		}
	default:
		mySuperServer = new(MyServer)
	}
	http.Handle("/", mySuperServer)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
