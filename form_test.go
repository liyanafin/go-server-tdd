package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestFormHandler_ValidSubmission(t *testing.T) {
	form := url.Values{}
	form.Add("name", "John Doe")
	form.Add("email", "john@example.com")
	form.Add("subject", "Test Subject")
	form.Add("message", "This is a test message.")
	req := httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	formHandle(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Errorf("expected status 303 See Other, got %d", res.StatusCode)
	}
	location := res.Header.Get("Location")
	if !strings.Contains(location, "status=success") {
		t.Errorf("expected redirect to contain status=success, got %s", location)
	}
}

func TestFormHandler_MissingFields(t *testing.T) {
	// Missing name
	form := url.Values{}
	form.Add("email", "john@example.com")
	form.Add("subject", "Test Subject")
	form.Add("message", "This is a test message.")
	req := httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	formHandle(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Errorf("expected status 303 See Other for missing name, got %d", res.StatusCode)
	}
	location := res.Header.Get("Location")
	if !strings.Contains(location, "error=Missing+name") {
		t.Errorf("expected redirect to contain error=Missing+name, got %s", location)
	}

	// Missing email
	form = url.Values{}
	form.Add("name", "John Doe")
	form.Add("subject", "Test Subject")
	form.Add("message", "This is a test message.")
	req = httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	formHandle(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Errorf("expected status 303 See Other for missing email, got %d", res.StatusCode)
	}
	location = res.Header.Get("Location")
	if !strings.Contains(location, "error=Missing+email") {
		t.Errorf("expected redirect to contain error=Missing+email, got %s", location)
	}

	// Missing subject
	form = url.Values{}
	form.Add("name", "John Doe")
	form.Add("email", "john@example.com")
	form.Add("message", "This is a test message.")
	req = httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	formHandle(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Errorf("expected status 303 See Other for missing subject, got %d", res.StatusCode)
	}
	location = res.Header.Get("Location")
	if !strings.Contains(location, "error=Missing+subject") {
		t.Errorf("expected redirect to contain error=Missing+subject, got %s", location)
	}

	// Missing message
	form = url.Values{}
	form.Add("name", "John Doe")
	form.Add("email", "john@example.com")
	form.Add("subject", "Test Subject")
	req = httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	formHandle(w, req)
	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Errorf("expected status 303 See Other for missing message, got %d", res.StatusCode)
	}
	location = res.Header.Get("Location")
	if !strings.Contains(location, "error=Missing+message") {
		t.Errorf("expected redirect to contain error=Missing+message, got %s", location)
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
