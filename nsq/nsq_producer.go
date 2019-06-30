package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)


func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("182.150.31.132:4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}