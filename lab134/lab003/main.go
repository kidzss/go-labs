package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//读取数据，用|隔开
	conf, err := ioutil.ReadFile("zk.conf")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	conns := strings.Split(string(conf), "|")
	conn, ech, err := zk.Connect(conns, time.Second)
	if err != nil {
		log.Fatalf("zk.Connect error:%v", err)
	}
	go func() {
		for {
			select {
			case e := <-ech:
				log.Printf("get Event :%+v", e)
			}
		}
	}()

	_ = conn

	<-sigs
	log.Println("program end")
}
