package server


import (
	"os"

	log "github.com/AKovalevich/scrabbler/log/logrus"
	"github.com/AKovalevich/scrabbler/config"
	"sync"
	"net/http"
	"context"
	"time"
	"os/signal"
)

// Server is the reverse-proxy/load-balancer engine
type Server struct {
	mainConfiguration *config.ScrabblerConfiguration
	signals							chan os.Signal
	stopChan						chan bool
	mainHttpServer					*http.Server
	webUiHttpServer					*http.Server
}

func NewServer(config *config.ScrabblerConfiguration) Server {
	var server Server
	server.mainConfiguration = config

	log.Do.Info("Configure signals and listeners...")
	// Configure signals
	server.configureSignals()
	go server.listenSignals()

	log.Do.Info("Configure HTTP services...")
	// Configure main scrabbler HTTP server
	server.configureMainHttpServer()
	server.configureWebUiHttpServer()

	return server
}

func (server *Server) Serve() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		go server.runMainServer()
		go server.runWebUiServer()
	}()
	wg.Wait()
}

func (server * Server) Close() {
	// Close Web UI HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	if err := server.webUiHttpServer.Shutdown(ctx); err != nil {
		log.Do.Errorf("Error: %v\n", err)
	} else {
		log.Do.Info("Web server is stopped")
	}

	if err := server.webUiHttpServer.Shutdown(ctx); err != nil {
		log.Do.Errorf("Error: %v\n", err)
	} else {
		log.Do.Info("Main server is stopped")
	}

	signal.Stop(server.signals)
	close(server.signals)
	//close(server.stopChan)
	cancel()
}