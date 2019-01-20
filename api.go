package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	ch := make(chan map[string]interface{})
	id := strings.TrimPrefix(r.URL.Path, "/api/stats/")
	if id == "" {
		return
	}
	fmt.Println(id)
	go func() {
		defer func() {
			close(ch)
		}()
		url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%v?api_key=%v", id, key)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		var data map[string]interface{}

		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			fmt.Println(err)
		}
		if _, ok := data["id"]; !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ch <- data
	}()
	v := <-ch
	if _, ok := v["id"]; !ok {
		return
	}
	fmt.Println(v)
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/league/v4/positions/by-summoner/%v?api_key=%v", v["id"], key)
	rz, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(rz.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(body)
}
