# Noa

## a network framework

## Example

### custom handler

```go
type TestChanelHandler struct {
}

func (handler TestChanelHandler) OnMessage(chanel *io.Chanel, data *bytes.Buffer) {
log.Println(string(data.Bytes()))
}

func (handler TestChanelHandler) OnError(chanel *io.Chanel, err error) {
log.Println(err.Error())
}

func (handler TestChanelHandler) OnClose() {
log.Println("connection is closed")
}

```

### client

```go
client := io.NewClient(TestChanelHandler{}, codec.DefaultCodec)

_, chanel := client.Connect("127.0.0.1:10086")

chanel.WriteAndFlush(bytes.NewBuffer([]byte("data")))

```

### server

```go
server := io.NewServer(TestChanelHandler{}, codec.LengthSplitCodec{})

err := server.Listen(10086)

```



