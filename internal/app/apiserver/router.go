package apiserver

import (
	"io"
	"net/http"
)

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *ApiServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello World")
		//w.Write([]byte("Hello World"))
	}
}
