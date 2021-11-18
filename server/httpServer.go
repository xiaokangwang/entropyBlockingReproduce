package server

import (
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type HTTPStreamServer struct {
	net.Listener
}

func (s HTTPStreamServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(204)
}

func (s HTTPStreamServer) Serve() {
	log.Fatal(http.Serve(s.Listener, s))
}
