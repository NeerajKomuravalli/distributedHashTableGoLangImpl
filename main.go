package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/model"
	"github.com/gorilla/mux"
)

var HashTableData = model.Hashtable{}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	HashTableData.Lock.Lock()
	value, ifValueFound := HashTableData.Item[params["id"]]
	HashTableData.Lock.Unlock()
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
	HashTableData.Lock.Lock()
	_, ifValueFound := HashTableData.Item[userData.Id]
	if ifValueFound {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s you want to add is already present in the table, please use Update method to update any value", userData.Id),
		}
		json.NewEncoder(w).Encode(errorResp)
		HashTableData.Lock.Unlock()
		return
	}
	if HashTableData.Item == nil {
		HashTableData.Item = make(map[string]interface{})
	}
	HashTableData.Item[userData.Id] = userData.Value
	HashTableData.Lock.Unlock()
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
	HashTableData.Lock.Lock()
	_, ifValueFound := HashTableData.Item[userData.Id]
	if !ifValueFound {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s you want to update is not present in the table, please use Post method to add value first and then update", userData.Id),
		}
		json.NewEncoder(w).Encode(errorResp)
		HashTableData.Lock.Unlock()
		return
	}
	HashTableData.Item[userData.Id] = userData.Value
	HashTableData.Lock.Unlock()
	successResp := model.ResponseData{
		Success: true,
		Data:    &userData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	HashTableData.Lock.Lock()
	value, ifValueFound := HashTableData.Item[params["id"]]
	if !ifValueFound {
		errorResp := model.ResponseData{
			Success:      false,
			ErrorMessage: fmt.Sprintf("The id %s you want to delete is not present in the table", params["id"]),
		}
		json.NewEncoder(w).Encode(errorResp)
		HashTableData.Lock.Unlock()
		return
	}
	delete(HashTableData.Item, params["id"])
	HashTableData.Lock.Unlock()
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
