package main

import (
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

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/healthcheck", nil)

	app.healthcheckHandler(w, r)

	expectedBody := "Status: OK\nenvironment: development\nversion: 1.0.0\n"
	internal.Assert(t, w.Code, http.StatusOK)
	internal.Assert(t, w.Body.String(), expectedBody)
}
