package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// create a mock request
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	//create and record response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	//check status
	if s := rr.Code; s != http.StatusOK {
		t.Errorf("unexpected status code: got %v want %v", s, http.StatusOK)
	}

	//check body
	expected := `{"alive": true}`
	if v := rr.Body.String(); v != expected {
		t.Errorf("unexpected body: got %v want %v", v, expected)
	}
}
