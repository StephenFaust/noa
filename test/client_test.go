package main

import (
	"bytes"
	"log"
	"noa/codec"
	"noa/io"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	client := io.NewClient(TestChanelHandler{}, codec.DefaultCodec)
	_, chanel := client.Connect("127.0.0.1:10086")

	for {
		if err := chanel.WriteAndFlush(bytes.NewBuffer([]byte("1"))); err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second * 1)
	}

}
