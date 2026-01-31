package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestFormHandler_ValidSubmission(t *testing.T) {
	form := url.Values{}
	form.Add("name", "John Doe")
	form.Add("address", "123 Main St")
	req := httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	formHandle(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "Thank you") {
		t.Errorf("expected response to contain 'Thank you', got %s", string(body))
	}
}

func TestFormHandler_MissingName(t *testing.T) {
	form := url.Values{}
	form.Add("address", "123 Main St")
	req := httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	formHandle(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400 Bad Request, got %d", res.StatusCode)
	}
}

func TestFormHandler_GetMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/form", nil)
	w := httptest.NewRecorder()

	formHandle(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405 Method Not Allowed, got %d", res.StatusCode)
	}
}
