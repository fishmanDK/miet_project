package app

import "net/http"

func (a *app) createHttpServer() *http.Server {
	server := &http.Server{
		Addr:         a.cfg.HTTP.Port,
		Handler:      a.gin,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
	}

	return server
}