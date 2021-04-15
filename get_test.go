package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func AddParametersToRequest(req *http.Request, parameter string) *http.Request {
	vars := map[string]string{
		"id": parameter,
	}
	return mux.SetURLVars(req, vars)
}

func TestGetSampleDataSuccess(t *testing.T) {
	PostSampleData(t, `{"id":"112","value":112}`)

	req, err := http.NewRequest("GET", "/Get", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = AddParametersToRequest(req, "112")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":true,"data":{"id":"112","value":112}}`
	ResponseRecorderComparisonWithExpectedOut(t, rr, expected)
}

func TestGetSampleDataFailure(t *testing.T) {
	req, err := http.NewRequest("GET", "/Get", nil)
	vars := map[string]string{
		"id": "113",
	}
	req = mux.SetURLVars(req, vars)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":false,"error":"The id 113 is not found in the table"}`
	ResponseRecorderComparisonWithExpectedOut(t, rr, expected)
}
