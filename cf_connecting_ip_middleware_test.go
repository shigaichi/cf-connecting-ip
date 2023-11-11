package cfconnectingip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCfConnectingIPMiddleware(t *testing.T) {
	// Mock request with X-Forwarded-For and CF-Connecting-IP headers
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Forwarded-For", "1.2.3.4")
	req.Header.Set("CF-Connecting-IP", "5.6.7.8")

	// Recorder to capture the response
	recorder := httptest.NewRecorder()

	// Create a handler that uses the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Check if RemoteAddr was set correctly
		if r.RemoteAddr != "5.6.7.8" {
			t.Errorf("RemoteAddr not set correctly by middleware, got %s, want %s", r.RemoteAddr, "5.6.7.8")
		}
	})

	// Apply the middleware
	testHandler := SetRemoteAddr(handler)

	// Serve HTTP request to the recorder
	testHandler.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
