package probe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

type HTTPProbe struct {
	URL         string
	DialContext func(ctx context.Context, network string, addr string) (net.Conn, error)
}

func (p HTTPProbe) Probe() {
	transport := http.Transport{DialContext: p.DialContext, DisableKeepAlives: true}
	req, err := http.NewRequest("GET", p.URL, nil)
	if err != nil {
		panic(err)
	}
	_, err = transport.RoundTrip(req)
	if err != nil {
		log.WithField("action", "http").WithError(err).Info("unable to GET")
	}
	time.Sleep(time.Millisecond * 10)
}
