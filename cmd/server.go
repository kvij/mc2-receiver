package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api v1"))
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", version)
	api.HandleFunc("", version)

	api.HandleFunc("/apps", getAppsList).Methods(http.MethodGet)
	api.HandleFunc("/app/{id}", getAppInfo).Methods(http.MethodGet)
	api.HandleFunc("/app/{id}/info", getAppInfo).Methods(http.MethodGet)
	api.HandleFunc("/app/{id}", startApp).Methods(http.MethodPost)
	api.HandleFunc("/app/{id}/start", startApp).Methods(http.MethodPost, http.MethodPut)
	api.HandleFunc("/app/{id}", stopApp).Methods(http.MethodDelete)
	api.HandleFunc("/app/{id}/stop", stopApp).Methods(http.MethodDelete, http.MethodPost, http.MethodPut)
	api.HandleFunc("/app/{id}/kill", killApp).Methods(http.MethodDelete, http.MethodPost, http.MethodPut)
	log.Fatalln(http.ListenAndServe(":8888", r))
}
