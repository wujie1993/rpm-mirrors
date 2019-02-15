package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wujie1993/rpm-mirrors/core"
)

func SyncMirrors() error {
	conf := core.GetConf()
	for {
		// run mirror sync jobs
		for _, mirror := range conf.Mirrors {
			mirrorJob := &core.MirrorJob{
				Mirror:  mirror,
				Storage: conf.Storage,
			}
			if err := mirrorJob.Start(); err != nil {
				log.Println(err)
			}
			log.Println("succeed to sync " + mirrorJob.Mirror.Source)
		}
		log.Println("waiting for", conf.SyncInterval, "second and then start the next round ")
		// waiting for next round
		time.Sleep(time.Second * time.Duration(conf.SyncInterval))
	}
}

func main() {
	conf, err := core.LoadConf(core.DefaultConfPath)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	datas, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Println("Load config")
	log.Println(string(datas))

	go SyncMirrors()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case killSignal := <-interrupt:
			log.Println("Got signal: ", killSignal)
			return
		}
	}
}
