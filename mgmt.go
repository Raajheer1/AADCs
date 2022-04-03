package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Aircrafts struct {
	Aircrafts []Aircraft `json:"pilots"`
}

type Aircraft struct {
	CID        uint       `json:"cid"`
	Callsign   string     `json:"callsign"`
	Latitude   float32    `json:"latitude"`
	Longitude  float32    `json:"longitude"`
	Altitude   uint       `json:"altitude"`
	Groudspeed uint       `json:"groudspeed"`
	Heading    uint       `json:"heading"`
	Flightplan FlightPlan `json:"flight_plan"`
	Status     string
	Distance   float64 `json:"distance"`
}

type FlightPlan struct {
	Rules     string `json:"flight_rules"`
	Aircraft  string `json:"aircraft_faa"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Alternate string `json:"alternate"`
	TAS       uint   `json:"cruise_tas"`
	Altitude  uint   `json:"altitude"`
	Route     string `json:"route"`
	Routeprse string
}

func main() {
	fmt.Println("Starting State Management Server...")

}

func initialize() {
	vatsimData, err := http.Get("https://status.vatsim.net/status.json")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer vatsimData.Body.Close()
	contents, err := ioutil.ReadAll(vatsimData.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var aircraft Aircrafts
	json.Unmarshal(contents, &aircraft)
}
