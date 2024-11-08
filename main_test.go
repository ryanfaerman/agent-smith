package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExecuteHandlerPing(t *testing.T) {
	// Create a test HTTP server with the handler
	handler := http.NewServeMux()
	handler.HandleFunc("/execute", handleCommand(
		NewCommander(),
	))

	// Test data: CommandRequest to simulate a ping
	requestData := CommandRequest{
		Type:    "ping",
		Payload: "localhost",
	}
	body, err := json.Marshal(requestData)
	if err != nil {
		t.Fatalf("Failed to marshal request data: %v", err)
	}

	// Create a request to send to the handler
	req := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler with the request and recorder
	handler.ServeHTTP(rr, req)

	// Check the status code is OK (200)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	// Check the response body is correct (valid JSON response)
	var response CommandResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Validate the response based on the request data
	if !response.Success {
		t.Errorf("Expected success true, but got false. Error: %v", response.Error)
	}
	var pingResult PingResult
	dataBytes, err := json.Marshal(response.Data) // Marshal the Data field to handle it properly
	if err != nil {
		t.Fatalf("Failed to marshal Data: %v", err)
	}
	err = json.Unmarshal(dataBytes, &pingResult) // Unmarshal it to a PingResult
	if err != nil {
		t.Fatalf("Failed to unmarshal Data as PingResult: %v", err)
	}

	// Validate the ping result
	if !pingResult.Successful {
		t.Errorf("Expected successful ping, but got failure")
	}

	// Validate the time is a positive duration (check if ping was successful and has a reasonable time)
	if pingResult.Time <= 0 {
		t.Errorf("Expected a positive ping time, but got %v", pingResult.Time)
	}
}

func TestExecuteHandlerSysinfo(t *testing.T) {
	// Create a test HTTP server with the handler
	handler := http.NewServeMux()
	handler.HandleFunc("/execute", handleCommand(
		NewCommander(),
	))

	// Test data: CommandRequest to simulate a ping
	requestData := CommandRequest{
		Type: "sysinfo",
	}
	body, err := json.Marshal(requestData)
	if err != nil {
		t.Fatalf("Failed to marshal request data: %v", err)
	}

	// Create a request to send to the handler
	req := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler with the request and recorder
	handler.ServeHTTP(rr, req)

	// Check the status code is OK (200)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	// Check the response body is correct (valid JSON response)
	var response CommandResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Validate the response based on the request data
	if !response.Success {
		t.Errorf("Expected success true, but got false. Error: %v", response.Error)
	}
	var info SystemInfo
	dataBytes, err := json.Marshal(response.Data) // Marshal the Data field to handle it properly
	if err != nil {
		t.Fatalf("Failed to marshal Data: %v", err)
	}
	err = json.Unmarshal(dataBytes, &info) // Unmarshal it to a PingResult
	if err != nil {
		t.Fatalf("Failed to unmarshal Data as PingResult: %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.IPAddress == "" {
		t.Error("Expected IP address to be non-empty")
	}
}
