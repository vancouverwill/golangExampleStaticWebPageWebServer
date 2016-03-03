package main

import (
	//	"bytes"
	//	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePageGet(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:4000/", nil)
	if err != nil {
		t.Error("GET did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	Handler(w, req)

	if w.Code != 200 && w.Code != 202 {
		t.Error("GET did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	if !strings.Contains(w.Body.String(), "home") {
		t.Error("GET homepage did not return correct content ", w.Body.String(), " I expected to see some reference to home")
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}

func TestAboutUsGet(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:4000/aboutus/", nil)
	if err != nil {
		t.Error("GET did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	AboutUsHandler(w, req)

	if w.Code != 200 && w.Code != 202 {
		t.Error("GET did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	if !strings.Contains(w.Body.String(), "about") {
		t.Error("GET homepage did not return correct content ", w.Body.String(), " I expected to see some reference to \"about\"")
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}
