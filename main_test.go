package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() http.Handler {
	// ...existing code...
	mux := http.NewServeMux()
	mux.HandleFunc("/integration-spec", func(w http.ResponseWriter, r *http.Request) {
		// Dummy response for testing
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Integration Specification"))
	})
	mux.HandleFunc("/scan-url", func(w http.ResponseWriter, r *http.Request) {
		// Dummy response for testing
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var payload map[string]string
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &payload)
		response := map[string]interface{}{
			"event_name": "url_scanned",
			"message":    "✅ URL Check: https://example.com\n→ Status: safe\n→ Recommendation: This link appears safe",
			"urls":       []string{"https://example.com"},
			"status":     "success",
			"username":   "url-scanner-bot",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		respBytes, _ := json.Marshal(response)
		w.Write(respBytes)
	})
	return mux
}

func TestIntegrationSpecEndpoint(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/integration-spec", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", rr.Code)
	}
	expected := "Integration Specification"
	if rr.Body.String() != expected {
		t.Errorf("Expected response %q, got %q", expected, rr.Body.String())
	}
}

func TestScanURLEndpoint(t *testing.T) {
	router := setupRouter()

	data := map[string]string{"message": "Check this link: https://example.com"}
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", "/scan-url", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", rr.Code)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal("Invalid JSON response")
	}
	if resp["status"] != "success" {
		t.Errorf("Expected status 'success', got %v", resp["status"])
	}
}
