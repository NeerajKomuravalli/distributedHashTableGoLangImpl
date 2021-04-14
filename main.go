package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/model"
	"github.com/gorilla/mux"
)

var hashTableData = model.Hashtable{}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	hashTableData.Lock.Lock()
	value, ifValueFound := hashTableData.Item[params["id"]]
	hashTableData.Lock.Unlock()
	if !ifValueFound {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s is not found in the table", params["id"]),
		}
		json.NewEncoder(w).Encode(errorResp)
		return
	}
	userData := model.UserSentData{
		Id:    params["id"],
		Value: value,
	}
	successResp := model.ResponseData{
		Success: true,
		Data:    &userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func postData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userData := model.UserSentData{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: err.Error(),
		}
		json.NewEncoder(w).Encode(errorResp)
		return
	}
	hashTableData.Lock.Lock()
	_, ifValueFound := hashTableData.Item[userData.Id]
	if ifValueFound {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s you want to add is already present in the table, please use Update method to update any value", userData.Id),
		}
		json.NewEncoder(w).Encode(errorResp)
		hashTableData.Lock.Unlock()
		return
	}
	if hashTableData.Item == nil {
		hashTableData.Item = make(map[string]interface{})
	}
	hashTableData.Item[userData.Id] = userData.Value
	hashTableData.Lock.Unlock()
	successResp := model.ResponseData{
		Success: true,
		Data:    &userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func updateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userData := model.UserSentData{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: err.Error(),
		}
		json.NewEncoder(w).Encode(errorResp)
		return
	}
	hashTableData.Lock.Lock()
	_, ifValueFound := hashTableData.Item[userData.Id]
	if !ifValueFound {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s you want to update is not present in the table, please use Post method to add value first and then update", userData.Id),
		}
		json.NewEncoder(w).Encode(errorResp)
		hashTableData.Lock.Unlock()
		return
	}
	hashTableData.Item[userData.Id] = userData.Value
	hashTableData.Lock.Unlock()
	successResp := model.ResponseData{
		Success: true,
		Data:    &userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	hashTableData.Lock.Lock()
	value, ifValueFound := hashTableData.Item[params["id"]]
	if !ifValueFound {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s you want to delete is not present in the table", params["id"]),
		}
		json.NewEncoder(w).Encode(errorResp)
		hashTableData.Lock.Unlock()
		return
	}
	delete(hashTableData.Item, params["id"])
	hashTableData.Lock.Unlock()
	userData := model.UserSentData{
		Id:    params["id"],
		Value: value,
	}
	successResp := model.ResponseData{
		Success: true,
		Data:    &userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Get/{id}", getData).Methods("GET")
	router.HandleFunc("/Post/", postData).Methods("POST")
	router.HandleFunc("/Update/", updateData).Methods("PUT")
	router.HandleFunc("/Delete/{id}", deleteData).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
