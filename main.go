package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FlightPlan struct {
	Rules 		string 		`json:"flight_rules"`
	Aircraft 	string 		`json:"aircraft_faa"`
	Departure 	string 		`json:"departure"`
	Arrival 	string 		`json:"arrival"`
	Alternate 	string 		`json:"alternate"`
	TAS 		uint 		`json:"cruise_tas"`
	Altitude 	uint 		`json:"altitude"`
	Deptime 	uint 		`json:"deptime"`
	ERT 		uint 		`json:"enroute_time"`
	Fuel 		uint 		`json:"fuel_time"`
	Remarks 	string 		`json:"remarks"`
	Route 		string 		`json:"route"`
}

type Aircraft struct {
	CID			uint 		`json:"cid"`
	Callsign 	string 		`json:"callsign"`
	Latitude	float32		`json:"latitude"`
	Longitude	float32		`json:"longitude"`
	Altitude 	uint 		`json:"altitude"`
	Groudspeed 	uint 		`json:"groudspeed"`
	Heading 	uint 		`json:"heading"`
	Flightplan 	FlightPlan 	`json:"flight_plan"`
	Status 		string 		`json:"status"`
	EDT 		uint 		`json:"edt"`
	DT	 		uint 		`json:"dt"`
	Distance	uint		`json:"distance"`
	Arrival 	uint 		`json:"arrival"`
}

type Aircrafts struct {
	Aircrafts 	[]Aircraft 	`json:"pilots"`
}

func main() {
	resp, err := http.Get("https://data.vatsim.net/v3/vatsim-data.json")

	if(err != nil){
		fmt.Printf("HTTP GET error: ", err)
		return
	}else{
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if(err != nil){
			fmt.Printf("%s", err)
			return
		}

		var aircraft Aircrafts
		json.Unmarshal(contents, &aircraft)

		for i := 0; i < len(aircraft.Aircrafts); i++ {
			fmt.Println("Callsign: " + aircraft.Aircrafts[i].Callsign)
			fmt.Println("Departure: " + aircraft.Aircrafts[i].Flightplan.Departure)
			fmt.Println("Arrival: " + aircraft.Aircrafts[i].Flightplan.Arrival)
			fmt.Println("Altitude: " + fmt.Sprint(aircraft.Aircrafts[i].Altitude))
			fmt.Println("\n")

			if aircraft.Aircrafts[i].Status != "" {
				if aircraft.Aircrafts[i].Distance < 10 {
					aircraft.Aircrafts[i].Status = "Arrived"
				}else if aircraft.Aircrafts[i].Groudspeed > 50 {
					aircraft.Aircrafts[i].Status = "Flight Active"
				}else if aircraft.Aircrafts[i].EDT < CurrentTime {
					// TO DO ^^^
					aircraft.Aircrafts[i].Status = "Past Dept Time"
				}else{
					aircraft.Aircrafts[i].Status = "Departing"
				}
			}
		}
	}
}
