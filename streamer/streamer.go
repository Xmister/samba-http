package streamer

import (
	"os"
	"net/http"
	"strings"
	"context"

)

var srv http.Server


type pathTranslator struct {
}


func (k *pathTranslator) Open(path string) (file http.File, err error) {
	if strings.Contains(path[1:], "/") {
		file, err = NewSmbHandler().Open("smb://" + path[1:])
	}
	if file == nil || err != nil {
		err = os.ErrNotExist
	}
	return
}


//export StartServer
func StartServer() {
	srv.Addr="127.0.0.1:8080"
	srv.Handler=http.FileServer(&pathTranslator{})
	go func() {srv.ListenAndServe()}()
}

//export StopServer
func StopServer() {
	srv.Shutdown(context.Background())
}