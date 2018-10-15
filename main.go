package main

import (
	"net/http"
	"time"

	log "github.com/cihub/seelog"
)

func main() {
	initialize()
	log.Debug("Start main")
	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Server.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/info", serverInfo)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	/*
		// defined in route_auth.go
		mux.HandleFunc("/logout", logout)
		mux.HandleFunc("/signup_account", signupAccount)
		mux.HandleFunc("/authenticate", authenticate)
	*/
	// starting up the server
	server := &http.Server{
		Addr:           config.Server.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.Server.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.Server.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("Start Server" + config.Server.Address)
	server.ListenAndServe()
}
