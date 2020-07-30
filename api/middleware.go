package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/core"
	inlog "github.com/bots-house/birzzha/pkg/log"
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

func (h *Handler) wrapMiddlewareLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = inlog.WithLogger(ctx, h.Logger)

		// check if request id header is present and add it to context.
		if reqID := r.Header.Get("X-Request-ID"); reqID != "" {
			ctx = inlog.WithPrefix(ctx, "request_id", reqID)
		}

		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}

func (h *Handler) wrapMiddlewareRecovery(handler http.Handler) http.Handler {
	ctx := context.Background()
	ctx = inlog.WithLogger(ctx, h.Logger)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				inlog.Error(ctx, "something went wrong, recovery", "URL", r.URL, "method", r.Method, "error", err)

				fmt.Println(string(debug.Stack()))

				w.Header().Set("Content-Type", "appilication/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"code": "internal_error"}`))
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
