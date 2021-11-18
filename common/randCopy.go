package common

import (
	log "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"time"
)

func RandData(writer io.Writer) error {
	for {
		length := rand.Intn(6553)
		waitTime := rand.Intn(1000)
		_, err := io.CopyN(writer, rand.New(rand.NewSource(time.Now().UnixNano())), int64(length))
		if err != nil {
			log.WithField("type", "write").
				WithField("time", time.Now().String()).
				WithField("result", "failure").Infoln("Unable to finish write")
			return err
		}
		log.WithField("type", "write").
			WithField("time", time.Now().String()).
			WithField("result", "success").
			WithField("amount", length).
			WithField("sleep", waitTime).
			Infoln("data written to remote")
		time.Sleep(time.Millisecond * time.Duration(waitTime))
	}
}
