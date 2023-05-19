package main

import (
	"bytes"
	"github.com/StephenFaust/noa/codec"
	"github.com/StephenFaust/noa/io"
	"log"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	client := io.NewClient(TestChanelHandler{}, codec.DefaultCodec)
	_, chanel := client.Connect("127.0.0.1:10086")

	//client2 := io.NewClient(TestChanelHandler{}, codec.DefaultCodec)
	//_, chanel2 := client2.Connect("127.0.0.1:10086")
	for {
		if err := chanel.WriteAndFlush(bytes.NewBuffer([]byte("1"))); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 5)
		//chanel2.WriteAndFlush(bytes.NewBuffer([]byte("2")))
		//
		//time.Sleep(time.Second * 5)
		//
		//chanel.WriteAndFlush(bytes.NewBuffer([]byte("1")))
		//time.Sleep(time.Second * 5)
		//chanel2.WriteAndFlush(bytes.NewBuffer([]byte("2")))
	}

}
