package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Coaster struct {
	Name         string `json:name`
	Manufacturer string `json:manufacturer`
	ID           string `json:id`
	InPark       string `json:inPark`
	Height       int    `json:height`
}

type coastersHandlers struct {
	sync.Mutex
	store map[string]Coaster
}

func (h *coastersHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (h *coastersHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	h.Lock()
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	h.Unlock()

	jsonByte, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

// curl localhost:8080/coasters -X POST -d '{"name": "Taron", "inPark" : "Phantasialand", "height": 30, "manufacturer": "Intamin"}' -H "Content-type: application/json"
func (h *coastersHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content-type 'application/json' but got '%s'", ct)))
		return
	}

	var coaster Coaster
	err = json.Unmarshal(bodyBytes, &coaster)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Lock()
	h.store[coaster.ID] = coaster
	defer h.Unlock()
}

func newCoasterHandler() *coastersHandlers {
	return &coastersHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				Name:         "Fury 325",
				Height:       99,
				InPark:       "Carowinds",
				Manufacturer: "B&M",
			},
		},
	}
}

func main() {
	coastersHandlers := newCoasterHandler()
	http.HandleFunc("/coasters", coastersHandlers.coasters)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
