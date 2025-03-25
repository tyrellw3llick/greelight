package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	internal "github.com/tyrellw3llick/greenlight/internal/assert"
)

func TestHealthCheck(t *testing.T) {
	app := &application{
		config: config{
			port: 4000,
			env:  "development",
		},
	}

	expectedData := map[string]string{
		"status":      "OK",
		"environment": "development",
		"version":     "1.0.0",
	}
	gotData := make(map[string]string)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/healthcheck", nil)

	app.healthcheckHandler(w, r)

	internal.Assert(t, w.Code, http.StatusOK)

	err := json.Unmarshal(w.Body.Bytes(), &gotData)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	internal.AssertMap(t, gotData, expectedData)
}
