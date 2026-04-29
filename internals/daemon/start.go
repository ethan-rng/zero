package daemon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server represents the daemon REST server
type Server struct {
	httpServer *http.Server
	local      bool
}

// NewServer creates a new daemon Server instance
func NewServer(port string, local bool) *Server {
	mux := http.NewServeMux()
	
	s := &Server{
		local: local,
	}

	// Register basic API routes
	mux.HandleFunc("/api/status", s.handleStatus)
	
	s.httpServer = &http.Server{
		Addr:    port,
		Handler: mux,
	}

	return s
}

// Start begins listening and serving requests. It blocks until the server stops.
func (s *Server) Start() error {
	log.Printf("Daemon server listening on %s (Local: %v)", s.httpServer.Addr, s.local)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

// Stop gracefully shuts down the server.
func (s *Server) Stop(ctx context.Context) error {
	log.Println("Shutting down daemon server...")
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]any{
		"status": "online",
		"local":  s.local,
		"time":   time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode status response: %v", err)
	}
}
