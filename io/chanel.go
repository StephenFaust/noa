package io

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/StephenFaust/noa/codec"
	"io"
	"net"
	"sync"
	"sync/atomic"
)

type Chanel struct {
	r       io.Reader
	w       io.Writer
	c       io.Closer
	cc      codec.Codec
	handler ChanelHandler
	isClose atomic.Bool
	locker  sync.Mutex
}

func (c *Chanel) WriteAndFlush(data *bytes.Buffer) (err error) {
	c.locker.Lock()
	defer c.locker.Unlock()
	err = c.cc.Encode(c.w, data)
	if err != nil {
		return err
	}
	err = c.w.(*bufio.Writer).Flush()
	return err
}

func (c *Chanel) IsActive() bool {
	return !c.isClose.Load()
}

func (c *Chanel) readyToRead() {
	go func() {
		for c.IsActive() {
			readData(c)
		}
	}()
}

func readData(c *Chanel) {
	defer catchError(c)
	data, err := c.cc.Decode(c.r)
	if err != nil {
		if _, ok := err.(net.Error); ok || err == io.EOF {
			// 处理连接异常错误
			c.Close()
		} else {
			// 处理其他错误
			c.handler.OnError(c, err)
		}
	} else {
		c.handler.OnMessage(c, data)
	}
}

func getChanel(conn net.Conn, cc codec.Codec, handler ChanelHandler) *Chanel {
	return &Chanel{r: bufio.NewReader(conn),
		w:       bufio.NewWriter(conn),
		c:       conn,
		cc:      cc,
		handler: handler,
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
