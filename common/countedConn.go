package common

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net"
	"sync/atomic"
	"time"
)

func NewCountedDialer(name string) *countedConnDialer {
	ccd := countedConnDialer{name: name}
	go ccd.checker()
	return &ccd
}

type countedConnDialer struct {
	readCount  int64
	writeCount int64
	name       string
}

func (d *countedConnDialer) DialContext(ctx context.Context, network string, addr string) (net.Conn, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return countedConnDialerConn{conn, d}, nil
}

func (d *countedConnDialer) checker() {
	lastread := int64(0)
	lastWrite := int64(0)
	for {
		thisReadCount := atomic.LoadInt64(&d.readCount)
		thisWriteCount := atomic.LoadInt64(&d.writeCount)
		readDelta := thisReadCount - lastread
		writeDelta := thisWriteCount - lastWrite
		lastread = thisReadCount
		lastWrite = thisWriteCount
		log.WithField("time", time.Now()).
			WithField("readDelta", readDelta).
			WithField("writeDelta", writeDelta).
			WithField("action", "typedCounter").
			WithField("name", d.name).Infoln("dialerTick")
		time.Sleep(time.Second)
	}
}

type countedConnDialerConn struct {
	net.Conn
	parent *countedConnDialer
}

func (c countedConnDialerConn) Read(p []byte) (n int, err error) {
	n, err = c.Conn.Read(p)
	atomic.AddInt64(&c.parent.readCount, int64(n))
	return
}

func (c countedConnDialerConn) Write(p []byte) (n int, err error) {
	n, err = c.Conn.Write(p)
	atomic.AddInt64(&c.parent.writeCount, int64(n))
	return
}
