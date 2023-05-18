package io

import (
	"fmt"
	"github.com/StephenMAOhjm/noa/codec"
	"net"
)

type Client struct {
	chanel  *Chanel
	cc      codec.Codec
	handler ChanelHandler
}

func NewClient(handler ChanelHandler, cc codec.Codec) *Client {
	client := new(Client)
	client.addHandler(handler)
	client.addCodec(cc)
	return client
}

func (c *Client) addHandler(handler ChanelHandler) {
	c.handler = handler
}

func (c *Client) addCodec(cc codec.Codec) {
	c.cc = cc
}

func (c *Client) Connect(address string) (error, *Chanel) {
	if c.handler == nil {
		return fmt.Errorf("no handler,please set"), nil
	}
	if c.cc == nil {
		return fmt.Errorf("no codec,please set"), nil
	}
	conn, err := net.Dial(TYPE, address)
	if err != nil {
		return err, nil
	}
	chanel := getChanel(conn, c.cc, c.handler)
	c.chanel = chanel
	c.handler.OnActive(chanel)
	c.chanel.readyToRead()
	return nil, chanel
}
