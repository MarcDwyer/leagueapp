package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/RiotSchmick?api_key=<key>

var key string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key = os.Getenv("KEY")
	fmt.Println(key)
}

// /lol/summoner/v4/summoners/by-name/{summonerName}
func main() {
	http.HandleFunc("/api/stats", stats)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func stats(w http.ResponseWriter, r *http.Request) {
	ch := make(chan map[string]interface{})
	go func() {
		defer func() {
			close(ch)
		}()
		url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/Santorin?api_key=%v", key)
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
			return
		}
		ch <- data
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
