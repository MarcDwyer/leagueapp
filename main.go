package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

// https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/RiotSchmick?api_key=<key>

var key string

func init() {
	fmt.Println(runtime.NumCPU())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key = os.Getenv("KEY")
	fmt.Println(key)
}

// /lol/summoner/v4/summoners/by-name/{summonerName}
func main() {
	http.HandleFunc("/api/stats/", Stats)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
