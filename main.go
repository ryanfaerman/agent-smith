package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
	Successful bool          `json:"successful"`
	Time       time.Duration `json:"time"`
}

type SystemInfo struct {
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ip_address"`
}

// CommandRequest represents the incoming JSON request
type CommandRequest struct {
	Type    string `json:"type"`    // "ping" or "sysinfo"
	Payload string `json:"payload"` // For ping, this is the host
}

// CommandResponse represents the outgoing JSON response
type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func main() {
	commander := NewCommander()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handleRequests(commander),
	}
	log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

// handleCommand is a http.HandlerFunc that processes incoming requests
// and sends back the response.
func handleCommand(cmdr Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Decode the JSON body into the CommandRequest struct
		var request CommandRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var response CommandResponse

		// Process the command based on the Type
		switch request.Type {
		case "ping":
			result, err := cmdr.Ping(request.Payload)

			response.Success = true
			response.Data = result
			if err != nil {
				response.Error = err.Error()
			}

		case "sysinfo":
			result, err := cmdr.GetSystemInfo()

			response.Success = true
			response.Data = result
			if err != nil {
				response.Error = err.Error()
			}

		default:
			response.Success = false
			response.Error = "Invalid command type"
		}

		// Set the Content-Type to application/json
		w.Header().Set("Content-Type", "application/json")

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to create response", http.StatusInternalServerError)
			return
		}

		// Send the response
		w.Write(responseJSON)
	}
}
