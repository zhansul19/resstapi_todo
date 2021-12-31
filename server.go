package todo

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(post string,handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + post,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}
func(s *Server)ShutDown(ctx context.Context)error{
	return s.httpServer.Shutdown(ctx)
}
