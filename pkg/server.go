package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	client *OsloBikesClient
	ctx    context.Context
	port   int
}

func NewServer(ctx context.Context, port int, osloBikesClient *OsloBikesClient) *Server {
	return &Server{
		client: osloBikesClient,
		ctx:    ctx,
		port:   port,
	}
}

func (server *Server) GetAvailableBikesAndDocks(w http.ResponseWriter, req *http.Request) {
	availableBikesDocksResponse, err := server.client.GetAvailableBikesAndDocks()
	if err != nil {
		log.Printf("Server error: %+v", err)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	js, err := json.Marshal(availableBikesDocksResponse)
	if err != nil {
		log.Printf("Server error: %+v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		log.Printf("Server error: %+v", err)
		return
	}
}

func (server *Server) Start() error {
	http.HandleFunc("/", server.GetAvailableBikesAndDocks)
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)
	if err != nil {
		return err
	}
	return nil
}
