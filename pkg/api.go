package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var host = "0.0.0.0"
var port = "8081"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	router := mux.NewRouter().StrictSlash(true)
	log.Info("Setting API on path prefi '/api'") 
	apiRouter := router.PathPrefix("/api").Subrouter()
	log.Info("Setting method GET /encode")                                           
	apiRouter.PathPrefix("/encode").HandlerFunc(GetEncode).Methods("GET")
	log.Info("Setting method GET /decode")
	apiRouter.PathPrefix("/decode").HandlerFunc(GetDecode).Methods("GET")
	log.Info("Setting method GET /healthcheck")
	apiRouter.PathPrefix("/healthcheck").HandlerFunc(GetHealthCheck).Methods("GET") 
	log.Info("Server is running on " + host + ":" + port)                          
	http.ListenAndServe(host+":"+port, router)
}

func GetEncode(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("string")
	log.Info("Request to Encode "+query )
	payload := Encode(r.URL.Query().Get("string"))
	response, _ := json.Marshal(payload)
	log.Debug("Response: " + string(response))
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetDecode(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("string")
	log.Info("Request to Encode "+query )
	payload, err := Decode(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	response, _ := json.Marshal(payload)
	log.Debug("Response: " + string(response))
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}