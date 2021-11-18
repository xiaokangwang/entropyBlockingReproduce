package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/entropyBlockingReproduce/common"
	"net"
)

type RandomStreamServer struct {
	net.Listener
}

func (s RandomStreamServer) Serve() {
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.serveConn(conn)
	}
}

func (s RandomStreamServer) serveConn(conn net.Conn) {
	go common.RandData(conn)
	common.ReadAll(conn)
}
