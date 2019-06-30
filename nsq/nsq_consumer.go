package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(string(message.Body))
		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQLookupd("182.150.31.132:4161")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}

