package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (s *Server) EventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a new message channel for this client
	messageChan := make(chan string)
	s.clients[messageChan] = true
	slog.Info("Client connected")

	// Ensure the channel is closed when the client disconnects
	defer func() {
		delete(s.clients, messageChan)
		close(messageChan)
		slog.Info("Client disconnected")
	}()

	// Listen for messages and send to the client
	for msg := range messageChan {
		fmt.Fprintf(w, "data: %s\n\n", msg)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func (s *Server) HandleMessages(broadcast chan string) {
	for {
		msg := <-broadcast
		slog.Info("Broadcasting message", "message", msg, "clients", len(s.clients))
		for client := range s.clients {
			client <- msg
		}
	}
}
