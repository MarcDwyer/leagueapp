package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Payload struct {
	Summoner map[string]interface{} `json:"summoner"`
	Match    map[string]interface{} `json:"match"`
}

func Stats(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	id := strings.TrimPrefix(r.URL.Path, "/api/stats/")
	if id == "" {
		return
	}

	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%v?api_key=%v", id, key)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	if _, ok := data["id"]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url = fmt.Sprintf("https://na1.api.riotgames.com/lol/league/v4/positions/by-summoner/%v?api_key=%v", data["id"], key)
	rz, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	var payload Payload
	err = json.NewDecoder(rz.Body).Decode(&payload.Summoner)
	if err != nil {
		fmt.Println(err)
	}
	url = fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/%v?api_key=%v", data["accountId"], key)
}

///lol/match/v4/matchlists/by-account/{encryptedAccountId}

func getStats(url string, data map[string]interface{}) {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/%v?api_key=%v", data["accountId"], key)
}
