package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type OSMResponse struct {
	PlaceID   int    `json:"place_id"`
	PlaceName string `json:"display_name"`
	Lat       string `json:"lat"`
	Long      string `json:"lon"`
	Type      string `json:"type"`
}

const api_url = "https://nominatim.openstreetmap.org/search"

func PlaceIdentifier(params string) {
	resp, err := http.Get(api_url + params)
	if err != nil {
		log.Println("Error Get Response :", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error Read Response Body :", err)
		return
	}

	var responses []OSMResponse
	err = json.Unmarshal(respBody, &responses)
	if err != nil {
		log.Println("Error Unmarshal Body Data :", err)
		return
	}

	for _, loc := range responses {
		fmt.Printf("Place ID : %d\n", loc.PlaceID)
		fmt.Printf("Name : %s\n", loc.PlaceName)
		fmt.Printf("Latitude : %s\n", loc.Lat)
		fmt.Printf("Longtitude : %s\n", loc.Long)
		fmt.Printf("Type : %s\n", loc.Type)
		fmt.Println("------------------------")
	}
}
