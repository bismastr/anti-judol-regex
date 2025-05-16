package server

import "net/http"

type Server struct {
	*http.Server
}

func NewServer() *Server {

	srv := http.Server{
		Addr: ":8080",
	}

	return &Server{&srv}
}
