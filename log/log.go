package main

import (
	"flag"

	"github.com/appscode/go/log"
	"github.com/appscode/go/log/golog"
)

func main() {
	golog.InitLogs()
	defer golog.FlushLogs()

	flag.Parse()

	// log.Fatalln("0: fatal")
	log.Errorln("1: error")
	log.Warning("2: warning")
	log.Infoln("3: info")
	log.Debugln("4: debug")
}
