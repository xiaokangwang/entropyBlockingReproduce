package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/entropyBlockingReproduce/common"
	"github.com/xiaokangwang/entropyBlockingReproduce/probe"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}
func main() {
	RANDAddr, ok := os.LookupEnv("RANDADDR")
	if !ok {
		log.Fatalf("no RANDADDR environment")
	}
	HTTPAddr, ok := os.LookupEnv("HTTPADDR")
	if !ok {
		log.Fatalf("no HTTPADDR environment")
	}
	randcd := common.NewCountedDialer("rand")
	randDialContext := randcd.DialContext
	httpcd := common.NewCountedDialer("http")
	httpDialContext := httpcd.DialContext

	go func() {
		httpprobe := probe.HTTPProbe{
			URL:         HTTPAddr,
			DialContext: httpDialContext,
		}
		for {
			httpprobe.Probe()
		}
	}()

	go func() {
		randp := probe.RandomProbe{
			Addr:        RANDAddr,
			DialContext: randDialContext,
		}
		for {
			randp.Probe()
		}
	}()

	<-make(chan struct{})
}
