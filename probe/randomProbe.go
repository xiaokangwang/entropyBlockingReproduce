package probe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/entropyBlockingReproduce/common"
	"net"
	"time"
)

type RandomProbe struct {
	Addr        string
	DialContext func(ctx context.Context, network string, addr string) (net.Conn, error)
}

func (p RandomProbe) Probe() {
	conn, err := p.DialContext(context.Background(), "tcp", p.Addr)
	if err != nil {
		log.WithField("action", "random").WithError(err).Errorln("unable to probe random")
		return
	}
	go func() {
		common.ReadAll(conn)
	}()
	go func() {
		common.RandData(conn)
	}()
	time.Sleep(time.Second * 4)
	conn.Close()
}
