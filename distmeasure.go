package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FlightPlan struct {
	Rules     string `json:"flight_rules"`
	Aircraft  string `json:"aircraft_faa"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Alternate string `json:"alternate"`
	TAS       string `json:"cruise_tas"`
	Altitude  string `json:"altitude"`
	Deptime   string `json:"deptime"`
	ERT       string `json:"enroute_time"`
	Fuel      string `json:"fuel_time"`
	Remarks   string `json:"remarks"`
	Route     string `json:"route"`
}

type Aircraft struct {
	CID         uint       `json:"cid"`
	Callsign    string     `json:"callsign"`
	Latitude    float32    `json:"latitude"`
	Longitude   float32    `json:"longitude"`
	Altitude    uint       `json:"altitude"`
	Groundspeed uint       `json:"groundspeed"`
	Heading     uint       `json:"heading"`
	Flightplan  FlightPlan `json:"flight_plan"`
	Status      string     `json:"status"`
	EDT         uint       `json:"edt"`
	DT          uint       `json:"dt"`
	Distance    float64    `json:"distance"`
	Arrival     uint       `json:"arrival"`
}

type Aircrafts struct {
	Aircrafts []Aircraft `json:"pilots"`
}

func main() {
	resp, err := http.Get("https://data.vatsim.net/v3/vatsim-data.json")
	//numPlanes := 0
	if err != nil {
		fmt.Printf("HTTP GET error: ", err)
		return
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		var aircraft Aircrafts
		json.Unmarshal(contents, &aircraft)

		fmt.Printf("%+v\n", aircraft.Aircrafts[0])
	}
}
