package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kvij/remotemc/app"
	"github.com/kvij/remotemc/xdg"
	"net/http"
)

func getAppsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(xdg.List())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getAppInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	if id, ok := p["id"]; ok {
		if a := app.New(id); a != nil {
			j, err := json.Marshal(a)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(j)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func startApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	if id, ok := p["id"]; ok {
		if a := app.New(id); a != nil {
			a.Start()
			a.Running = true
			j, err := json.Marshal(a)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(j)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func stopApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	if id, ok := p["id"]; ok {
		if a := app.New(id); a != nil {
			a.Stop()
			a.Running = false
			j, err := json.Marshal(a)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(j)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func killApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	if id, ok := p["id"]; ok {
		if a := app.New(id); a != nil {
			a.Kill()
			a.Running = false
			j, err := json.Marshal(a)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(j)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}