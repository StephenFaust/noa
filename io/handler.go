package io

import (
	"bytes"
)

type ChanelHandler interface {
	OnMessage(chanel *Chanel, data *bytes.Buffer)
	OnError(chanel *Chanel, err error)
	OnClose()
}
