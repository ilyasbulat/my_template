package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/ilyasbulat/rest_api/internal/config"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":8080"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, cfg *config.Config, opts ...Option) *Server {
	// c := cors.New(cors.Options{
	// 	AllowedMethods:     cfg.HTTP.CORS.AllowedMethods,
	// 	AllowedOrigins:     cfg.HTTP.CORS.AllowedOrigins,
	// 	AllowCredentials:   cfg.HTTP.CORS.AllowCredentials,
	// 	AllowedHeaders:     cfg.HTTP.CORS.AllowedHeaders,
	// 	OptionsPassthrough: cfg.HTTP.CORS.OptionsPassthrough,
	// 	ExposedHeaders:     cfg.HTTP.CORS.ExposedHeaders,
	// 	Debug:              cfg.HTTP.CORS.Debug,
	// })

	// handlerWithCORS := c.Handler(handler)

	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
