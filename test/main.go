package main

import (
	"strings"
)

func Print(args ...string) string {
	return strings.Join(args, ":")
}

type Server struct {
	addr string
}

func (s *Server) Addr() string {
	return s.addr
}

func (s *Server) getAddr() string {
	return s.addr
}

func main() {

}
