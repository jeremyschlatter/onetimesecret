package onetimesecret

import (
	"net/http"
	"os"

	"github.com/mholt/caddy/caddyhttp/httpserver"
)

type OneTimeSecret struct {
	Next     httpserver.Handler
	SiteRoot string
	Dirs     []string
}

func (s OneTimeSecret) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	status, err := s.Next.ServeHTTP(w, r)

	for _, dir := range s.Dirs {
		if !httpserver.Path(r.URL.Path).Matches(dir) {
			continue
		}
		removeErr := os.Remove(httpserver.SafePath(s.SiteRoot, r.URL.Path))
		if err == nil {
			err = removeErr
		}
		break
	}

	return status, err
}
