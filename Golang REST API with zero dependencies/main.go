package main

import (
	"encoding/json"
	"net/http"
)

type Coaster struct {
	Name         string `json:name`
	Manufacturer string `json:manufacturer`
	ID           string `json:id`
	InPark       string `json:inPark`
	Height       int    `json:height`
}

type coastersHandlers struct {
	store map[string]Coaster
}

func (h *coastersHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}

	jsonByte, err := json.Marshal(coasters)
	if err != nil {

	}

	w.Write(jsonByte)
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
	http.HandleFunc("/coasters", coastersHandlers.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
