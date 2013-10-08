package gaelog

import (
	"appengine"
	"log"
	"net/http"
	"os"
)

type gaeout struct {
	appengine.Context
}

func (self gaeout) Write(p []byte) (n int, err error) {
	self.Context.Infof(string(p))
	n = len(p)
	return
}

func Logger(r *http.Request) *log.Logger {
	if r != nil {
		return log.New(gaeout{appengine.NewContext(r)}, "", 0)
	} else {
		return log.New(os.Stdout, "", 0)
	}
}
