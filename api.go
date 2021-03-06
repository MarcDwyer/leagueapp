package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Payload struct {
	Summoner chan map[string]interface{} `json:"summoner, omitempty"`
	Match    chan map[string]interface{} `json:"match, omitempty"`
}
type Channel struct {
	Summoner map[string]interface{}
	Match    map[string]interface{}
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
		return
	}
	json.NewDecoder(res.Body).Decode(&data)
	if _, ok := data["id"]; !ok {
		fmt.Println(data)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var ch Payload
	// var super Channel

	go func() {
		for {
			select {
			case client, more := <-ch.Summoner:
				fmt.Println("test")
				fmt.Println(more)
				fmt.Println(client)
			}
		}
	}()
	fmt.Println(ch)
	str := fmt.Sprintf("/lol/league/v4/positions/by-summoner/%v?api_key=%v", data["id"], key)
	go GetStats("summonerInfo", str, data, &ch)

	str = fmt.Sprintf("/lol/match/v4/matchlists/by-account/%v?api_key=%v", data["accountId"], key)
	go GetStats("matches", str, data, &ch)
	fmt.Println(ch)
	fmt.Println(<-ch.Match)
}

///lol/match/v4/matchlists/by-account/{encryptedAccountId}

func GetStats(req string, url string, data map[string]interface{}, pointer *Payload) {
	fmt.Println("is this running?")
	fmt.Println(pointer)
	beg := fmt.Sprintf("https://na1.api.riotgames.com%v", url)
	switch req {
	case "summonerInfo":
		rz, err := http.Get(beg)
		if err != nil {
			fmt.Println(err)
			return
		}
		var usr []map[string]interface{}
		err = json.NewDecoder(rz.Body).Decode(&usr)
		if err != nil {
			fmt.Println("1")
			fmt.Println(err)
			return
		}
		pointer.Summoner <- usr[0]
		fmt.Println(<-pointer.Summoner)
	case "matches":
		rz, err := http.Get(beg)
		if err != nil {
			fmt.Println(err)
			return
		}
		var data map[string]interface{}
		err = json.NewDecoder(rz.Body).Decode(&data)
		if err != nil {
			fmt.Println("2")
			fmt.Println(err)
			return
		}
		pointer.Match <- data
	}
}
