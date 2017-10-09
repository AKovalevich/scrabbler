package server

import (
	"strconv"
	"net/http"
	"time"
	"fmt"

	log "github.com/AKovalevich/scrabbler/log/logrus"
)

func (server *Server) configureMainHttpServer() {
	s := http.NewServeMux()
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!\n")
	})
	server.mainHttpServer = &http.Server{
		Handler:      s,
		Addr: server.mainConfiguration.ServerHost + ":" + strconv.Itoa(server.mainConfiguration.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func (server *Server) configureWebUiHttpServer() {
	s := http.NewServeMux()
	s.HandleFunc("/ui", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!\n")
	})
	server.webUiHttpServer = &http.Server{
		Handler:      s,
		Addr: server.mainConfiguration.WebUIHost + ":" + strconv.Itoa(server.mainConfiguration.WebUIPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func (server *Server) runMainServer() {
	log.Do.Infof("Starting main scrabbler server on: %s", server.mainHttpServer.Addr)
	log.Do.Warn(server.mainHttpServer.ListenAndServe())
}

func (server *Server) runWebUiServer() {
	log.Do.Infof("Starting Web UI on: %s", server.webUiHttpServer.Addr)
	log.Do.Warn(server.webUiHttpServer.ListenAndServe())
}
