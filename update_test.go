package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func UpdateSampleData(t *testing.T, data string) *httptest.ResponseRecorder {
	var jsonStr = []byte(data)

	req, err := http.NewRequest("PUT", "/Update/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateData)
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

func TestUpdateDataSuccess(t *testing.T) {
	jsonStr := `{"id":"TestUpdateDataSuccess","value":"TestUpdateDataSuccess"}`
	PostSampleData(t, jsonStr)

	updatedJsonStr := `{"id":"TestUpdateDataSuccess","value":"TestUpdateDataSuccess"}`
	rr := UpdateSampleData(t, updatedJsonStr)
	expected := `{"success":true,"data":{"id":"TestUpdateDataSuccess","value":"TestUpdateDataSuccess"}}`
	ResponseRecorderComparisonWithExpectedOut(t, rr, expected)
}

func TestUpdateDataFailure(t *testing.T) {
	jsonStr := `{"id":"TestUpdateDataFailure","value":"TestUpdateDataFailure"}`
	rr := UpdateSampleData(t, jsonStr)
	expected := `{"success":false,"error":"The id TestUpdateDataFailure you want to update is not present in the table, please use Post method to add value first and then update"}`
	ResponseRecorderComparisonWithExpectedOut(t, rr, expected)
}
