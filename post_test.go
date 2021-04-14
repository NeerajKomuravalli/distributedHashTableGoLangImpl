package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostData(t *testing.T) {
	var jsonStr = []byte(`{"id":"111","value":111}`)

	req, err := http.NewRequest("GET", "/Post/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postData)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":true,"data":{"id":"111","value":111}}`
	// the strings that you read from STDIN has a trailing \n that's why we need to trip it
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
