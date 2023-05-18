package io

import (
	"bytes"
)

type ChanelHandler interface {
	OnActive(chanel *Chanel)
	OnMessage(chanel *Chanel, data *bytes.Buffer)
	OnError(chanel *Chanel, err error)
	OnClose()
}
