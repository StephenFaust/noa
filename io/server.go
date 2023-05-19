package io

import (
	"fmt"
	"github.com/StephenMAOhjm/noa/codec"
	"log"
	"net"
	"strconv"
)

const TYPE = "tcp"

type Server struct {
	cc      codec.Codec
	handler ChanelHandler
}

func NewServer(handler ChanelHandler, cc codec.Codec) *Server {
	server := new(Server)
	server.addHandler(handler)
	server.addCodec(cc)
	return server
}

func (s *Server) addHandler(handler ChanelHandler) {
	s.handler = handler
}

func (s *Server) addCodec(cc codec.Codec) {
	s.cc = cc
}

func (s *Server) Listen(port int) error {
	if s.handler == nil {
		return fmt.Errorf("no handler,please set")
	}
	if s.cc == nil {
		return fmt.Errorf("no codec,please set")
	}
	listen, err := net.Listen(TYPE, ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	handleConn(listen, s)
	return nil
}

func handleConn(listen net.Listener, s *Server) {
	go func() {
		for {
			chanel, err := accept(listen, s.cc, s.handler)
			if err != nil {
				log.Println("Error connection ", err)
				continue
			}
			chanel.readyToRead()
		}
	}()
}

func accept(listen net.Listener, cc codec.Codec, handler ChanelHandler) (*Chanel, error) {
	conn, err := listen.Accept()
	if err != nil {
		return nil, err
	}
	chanel := getChanel(conn, cc, handler)
	handler.OnActive(chanel)
	return chanel, nil
}
