package io

import (
	"net"
	"noa/codec"
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
	conn, err := net.Dial(TYPE, address)
	if err != nil {
		return err, nil
	}
	chanel := getChanel(conn, c.cc, c.handler)
	c.chanel = chanel
	c.chanel.readyToRead()
	return nil, chanel
}
