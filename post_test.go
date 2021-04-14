package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostSampleData(t *testing.T, data string) *httptest.ResponseRecorder {
	var jsonStr = []byte(data)

	req, err := http.NewRequest("POST", "/Post/", bytes.NewBuffer(jsonStr))
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
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	return rr
}

func TestPostDataSuccess(t *testing.T) {
	jsonStr := `{"id":"111","value":111}`
	rr := PostSampleData(t, jsonStr)
	expected := `{"success":true,"data":{"id":"111","value":111}}`
	// the strings that you read from STDIN has a trailing \n that's why we need to trip it
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostDataFailure(t *testing.T) {
	jsonStr := `{"id":"111","value":111}`
	PostSampleData(t, jsonStr)
	rr := PostSampleData(t, jsonStr)
	expected := `{"success":false,"error":"The id 111 you want to add is already present in the table, please use Update method to update any value"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
