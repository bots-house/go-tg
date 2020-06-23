package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/tg"
)

const filesPrefix = "/v1/tg/file/"

func newFileProxyWrapper(fp *tg.FileProxy) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if strings.HasPrefix(r.URL.Path, filesPrefix) {
				id := strings.TrimPrefix(r.URL.Path, filesPrefix)

				path, err := fp.Get(ctx, id)
				if err != nil {
					var body interface{}
					if err2, ok := errors.Cause(err).(*core.Error); ok {
						body = err2
					}

					if err := json.NewEncoder(w).Encode(body); err != nil {
						log.Printf("encode fail: %v", err)
					}
					return
				}

				http.ServeFile(w, r, path)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
