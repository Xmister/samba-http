package streamer

import (
	"github.com/Xmister/libdsm-go"
	"github.com/Xmister/libsmb2-go"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
)

type Dsm libdsm.Smb
type Smb2 libsmb2.Smb

type smbHandler struct {}

func NewSmbHandler() *smbHandler {
	return &smbHandler{}
}

func (s *smbHandler) ParseUrl(url string) (host string, user string, password string, share string, path string, err error) {
	var u *url2.URL
	if u, err =url2.Parse(url); err == nil {
		host = u.Host
		user = u.User.Username()
		if user == "" {
			user = "anonymous"
		}
		password, _ = u.User.Password()
		if u.Path == "" {
			share = u.Path
			path = u.Path
		} else if strings.Contains(u.Path[1:], "/") {
			spl := strings.SplitN(u.Path[1:], "/", 2)
			share = spl[0]
			path = spl[1]
		} else {
			share = u.Path[1:]
			path = ""
		}
	}
	return
}

func (s *smbHandler) Connect(url string) (fs http.FileSystem, err error) {
	var host, user, password, share string
	if host, user, password, share, _, err = s.ParseUrl(url); err == nil {
		smb := libsmb2.NewSmb()
		fs = (*Smb2)(smb)
		if err = smb.Connect(host, share, user, password); err != nil {
			smb := libdsm.NewSmb()
			fs = (*Dsm)(smb)
			err = smb.Connect(host, share, user, password)
		}
	}
	return
}

func (s *smbHandler) Open(url string) (file http.File, err error) {
	var smb http.FileSystem
	if smb, err = s.Connect(url); err == nil {
		var path string
		if _, _, _, _, path, err = s.ParseUrl(url); err == nil {
			file, err = smb.Open(path)
		}
	}
	return
}

func (s *Smb2) Open(path string) (file http.File, err error) {
	ss := (*libsmb2.Smb)(s)
	return ss.OpenFile(path, os.O_RDONLY)
}

func (s *Dsm) Open(path string) (file http.File, err error) {
	ss := (*libdsm.Smb)(s)
	return ss.OpenFile(path, os.O_RDONLY)
}

