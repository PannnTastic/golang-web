package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (l *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before LogMiddleware.ServeHTTP")
	l.Handler.ServeHTTP(w, r)
	fmt.Println("After LogMiddleware.ServeHTTP")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Executed")
		panic("Panic Handler")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	errorhandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: errorhandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type ErrorHandler struct {
	Handler http.Handler
}

func (e *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	e.Handler.ServeHTTP(w, r)
}
