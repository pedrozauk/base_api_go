package server

import (
	"fmt"
	"net/http"

	"github.com/pedropina/base_api/pkg/middleware"
)

type Server struct {
	mux *http.ServeMux
}

func (s *Server) InitServer(addr string) error {
	fmt.Printf("Escutando na porta: %s \n", addr)
	err := http.ListenAndServe(addr, s.mux)
	return err
}

func (s *Server) AddHandler(path string, handler func(http.ResponseWriter, *http.Request), acceptedContentsType ...string) {
	fmt.Println(len(acceptedContentsType))
	if len(acceptedContentsType) > 0 {
		handler = middleware.MiddleContentType(handler, acceptedContentsType)
	}
	s.mux.HandleFunc(path, handler)
}

func New() *Server {
	return &Server{
		http.NewServeMux(),
	}
}
