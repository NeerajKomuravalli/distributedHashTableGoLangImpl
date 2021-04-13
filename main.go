package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Hashtable struct {
	Item map[string]int
	Lock sync.RWMutex
}

type userSentData struct {
	Id    string `json:id`
	Value int    `json:value`
}

type responseData struct {
	Success      bool         `json:success`
	ErrorMessage string       `json:error,omitempty`
	Data         userSentData `json:data,omitempty`
}

var hashTableData = Hashtable{}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	value, ifValueFound := hashTableData.Item[params["id"]]
	if !ifValueFound {
		errorResp := responseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s is not found in the table", params["id"]),
		}
		json.NewEncoder(w).Encode(errorResp)
		return
	}
	userData := userSentData{
		Id:    params["id"],
		Value: value,
	}
	successResp := responseData{
		Success: true,
		Data:    userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func putData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userData := userSentData{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		errorResp := responseData{
			Success:      false,
			ErrorMessage: err.Error(),
		}
		json.NewEncoder(w).Encode(errorResp)
		return
	}
	if hashTableData.Item == nil {
		hashTableData.Item = make(map[string]int)
	}
	hashTableData.Item[userData.Id] = userData.Value
	successResp := responseData{
		Success: true,
		Data:    userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	value, ifValueFound := hashTableData.Item[params["id"]]
	if !ifValueFound {
		errorResp := responseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s you want to delete is not present in the table", params["id"]),
		}
		json.NewEncoder(w).Encode(errorResp)
		return
	}
	delete(hashTableData.Item, params["id"])
	userData := userSentData{
		Id:    params["id"],
		Value: value,
	}
	successResp := responseData{
		Success: true,
		Data:    userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Get/{id}", getData).Methods("GET")
	router.HandleFunc("/Put/", putData).Methods("POST")
	router.HandleFunc("/Delete/{id}", deleteData).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
