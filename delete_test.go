package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteSampleDataSuccess(t *testing.T) {
	PostSampleData(t, `{"id":"TestDeleteSampleDataSuccess","value":"TestDeleteSampleDataSuccess"}`)

	req, err := http.NewRequest("DELTE", "/Delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = AddParametersToRequest(req, "TestDeleteSampleDataSuccess")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":true,"data":{"id":"TestDeleteSampleDataSuccess","value":"TestDeleteSampleDataSuccess"}}`
	ResponseRecorderComparisonWithExpectedOut(t, rr, expected)
}

func TestDeleteSampleDataFailure(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/Delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = AddParametersToRequest(req, "TestDeleteSampleDataFailure")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteData)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":false,"error":"The id TestDeleteSampleDataFailure you want to delete is not present in the table"}`
	ResponseRecorderComparisonWithExpectedOut(t, rr, expected)
}
