package workers

import (
	"context"
	"net/http"
)

func Worker(ctx context.Context, requests chan *http.Request, handler http.Handler) {
	for {
		select {
		case req := <-requests:
			w := &responseWriter{}
			handler.ServeHTTP(w, req)
		case <-ctx.Done():
			return
		}
	}
}

type responseWriter struct {
	http.ResponseWriter
}
