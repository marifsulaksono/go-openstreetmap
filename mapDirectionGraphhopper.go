package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Instruction struct {
	Distance   float64 `json:"distance"`
	Heading    float64 `json:"heading"`
	Text       string  `json:"text"`
	Time       int     `json:"time"`
	StreetName string  `json:"street_name"`
}

type Path struct {
	Instructions []Instruction `json:"instructions"`
	BBox         any           `json:"bbox"`
}

type RouteResponse struct {
	Path []Path `json:"paths"`
}

const url = "https://graphhopper.com/api/1/route"

func GetMapDirectionUsingGraphhopperEngine(vehicle, startLat, startLong, endLat, endLong string) {
	params := fmt.Sprintf("?vehicle=%s&locale=en&key=LijBPDQGfu7Iiq80w3HzwB4RUDJbMbhs6BU0dEnn&elevation=false&"+
		"instructions=true&turn_costs=true&point=%s,%s&point=%s,%s",
		vehicle, startLat, startLong, endLat, endLong)

	resp, err := http.Get(url + params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var routeResponse RouteResponse
	err = json.Unmarshal(respBody, &routeResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("Bounding Box : %v\n", routeResponse.Path[0].BBox)
	for i, instruction := range routeResponse.Path[0].Instructions {
		time := time.Duration(instruction.Time) * time.Second
		fmt.Println("Step : ", i+1)
		fmt.Printf("Distance: %vm\n", instruction.Distance)
		fmt.Printf("Next Street: %v\n", instruction.StreetName)
		fmt.Printf("Instruction: %v\n", instruction.Text)
		fmt.Printf("Time: %v\n", time)
		fmt.Println("------------------------")
	}
}
