package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	jsonFile, err := os.Open("AircraftPerformance.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var Aircraft map[string]map[int]int
	json.Unmarshal(byteValue, &Aircraft)
	//Aircraft["B737"][36000] = 850
	AddAircraft(&Aircraft, "B737", 35000, 550)

	file, _ := json.MarshalIndent(Aircraft, "", " ")
	_ = ioutil.WriteFile("AircraftPerformance.json", file, 0644)
}

func AddAircraft(Aircraft *map[string]map[int]int, Type string, alt int, speed int) {

}
