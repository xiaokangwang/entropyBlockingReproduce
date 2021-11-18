package common

import (
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

func ReadAll(reader io.Reader) error {
	for {
		var buf [65536]byte
		n, err := reader.Read(buf[:])
		if err == nil {
			log.WithField("type", "read").
				WithField("time", time.Now().String()).
				WithField("result", "success").
				WithField("amount", n).
				Infoln("data read from remote")
		} else {
			log.WithField("type", "read").
				WithField("time", time.Now().String()).
				WithField("result", "failure").
				WithField("reason", err).
				Infoln("unable to data read from remote")
			return err
		}
	}
}
