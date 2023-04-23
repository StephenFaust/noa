package io

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/StephenMAOhjm/noa/codec"
	"io"
	"net"
	"sync/atomic"
)

type Chanel struct {
	r       io.Reader
	w       io.Writer
	c       io.Closer
	cc      codec.Codec
	handler ChanelHandler
	isClose atomic.Bool
}

func (c *Chanel) WriteAndFlush(data *bytes.Buffer) (err error) {
	err = c.cc.Encode(c.w, data)
	if err != nil {
		return err
	}
	return c.w.(*bufio.Writer).Flush()
}

func (c *Chanel) readyToRead() {
	go func() {
		for !c.isClose.Load() {
			if err := readData(c); err == io.EOF {
				c.isClose.Store(true)
			}
		}
	}()
}

func readData(c *Chanel) error {
	defer catchError(c)
	data, err := c.cc.Decode(c.r)
	if err != nil {
		c.handler.OnError(c, err)
		return err
	}
	c.handler.OnMessage(c, data)
	return err
}

func getChanel(conn net.Conn, cc codec.Codec, handler ChanelHandler) *Chanel {
	return &Chanel{bufio.NewReader(conn),
		bufio.NewWriter(conn),
		conn,
		cc,
		handler,
		atomic.Bool{},
	}
}

func catchError(c *Chanel) {
	if r := recover(); r != nil {
		c.handler.OnError(c, fmt.Errorf("runtime error: %v", r))
	}
}

func (c *Chanel) Close() {
	if c.isClose.Load() {
		return
	}
	defer catchError(c)
	c.isClose.Store(true)
	err := c.c.Close()
	if err != nil {
		c.handler.OnError(c, err)
	}
	c.handler.OnClose()
}
