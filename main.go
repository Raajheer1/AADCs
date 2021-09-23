package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type FlightPlan struct {
	Rules     string `json:"flight_rules"`
	Aircraft  string `json:"aircraft_faa"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Alternate string `json:"alternate"`
	TAS       uint   `json:"cruise_tas"`
	Altitude  uint   `json:"altitude"`
	Deptime   uint   `json:"deptime"`
	ERT       uint   `json:"enroute_time"`
	Fuel      uint   `json:"fuel_time"`
	Remarks   string `json:"remarks"`
	Route     string `json:"route"`
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
	Status     string     `json:"status"`
	EDT        uint       `json:"edt"`
	DT         uint       `json:"dt"`
	Distance   float64    `json:"distance"`
	Arrival    uint       `json:"arrival"`
}

type Aircrafts struct {
	Aircrafts []Aircraft `json:"pilots"`
}

func main() {
	start := time.Now()
	resp, err := http.Get("https://data.vatsim.net/v3/vatsim-data.json")
	numPlanes := 0
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

		for i := 0; i < len(aircraft.Aircrafts); i++ {
			if aircraft.Aircrafts[i].Flightplan.Route != "" && aircraft.Aircrafts[i].Flightplan.Rules == "I" && aircraft.Aircrafts[i].Flightplan.Arrival == "KDEN" {
				fmt.Println("Callsign: " + aircraft.Aircrafts[i].Callsign)
				fmt.Println("Route: " + aircraft.Aircrafts[i].Flightplan.Departure + " " + aircraft.Aircrafts[i].Flightplan.Route + " " + aircraft.Aircrafts[i].Flightplan.Arrival)
				aircraft.Aircrafts[i].Distance = Routedist(Routeparse(aircraft.Aircrafts[i].Flightplan.Departure + " " + aircraft.Aircrafts[i].Flightplan.Route + " " + aircraft.Aircrafts[i].Flightplan.Arrival))
				fmt.Printf("Distance: %.2f", aircraft.Aircrafts[i].Distance)
				fmt.Println("\n---")
				numPlanes++
			}
		}
	}

	duration := time.Since(start)
	fmt.Print("Counted: ")
	fmt.Println(numPlanes)
	fmt.Print("Runtime: ")
	fmt.Println(duration)

}
