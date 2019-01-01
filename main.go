package main

import (
	"os"
	"github.com/Xmister/samba-http/streamer"
)


func main() {
	signalChan := make(chan os.Signal, 1)
	streamer.StartServer()
	<- signalChan
	streamer.StopServer()
}
