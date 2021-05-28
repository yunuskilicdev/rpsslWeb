package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ChoicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(choices); err != nil {
		panic(err)
	}
}

func ChoiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	mod := GetRandomIndex()
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(choices[mod]); err != nil {
		panic(err)
	}
}

func GetRandomIndex() int {
	resp, err := http.Get("https://codechallenge.boohma.com/random")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var randomNumberResponse RandomNumberResponse
	json.NewDecoder(resp.Body).Decode(&randomNumberResponse)
	mod := randomNumberResponse.RandomNumber % len(choices)
	return mod
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var playRequest PlayRequest
	err := json.NewDecoder(r.Body).Decode(&playRequest)
	if err != nil  {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if playRequest.Player < 1 || playRequest.Player>len(choices) {
		http.Error(w, "Out of range choice", http.StatusBadRequest)
		return
	}
	randIndex := GetRandomIndex()
	computer := choices[randIndex]
	player := choices[playRequest.Player-1]
	result := results[player.Id][computer.Id]
	playResponse := PlayResponse{
		Results:  result,
		Player:   player.Id,
		Computer: computer.Id,
	}

	if err := json.NewEncoder(w).Encode(playResponse); err != nil {
		panic(err)
	}
}
