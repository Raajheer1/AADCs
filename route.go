package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Location struct {
	Lon		float32		`xml:"Lon,attr"`
	Lat		float32		`xml:"Lat,attr"`
}

type Fix struct {
	Name		string 		`xml:"ID,attr"`
	Location 	Location	`xml:"Location"`
}

type Fixes struct {
	Fixes []Fix `xml:"Waypoint"`
}

type Airway struct {
	Name	string
	Fixes 	[]string
}

var fixes Fixes = parseFIX("Waypoints.xml")
var airways []Airway = parseAWY("AWY.txt")

func main() {
	//FIXES
	//for i := 0; i < len(fixes.Fixes); i++ {
	//	fmt.Println("Name: " + fixes.Fixes[i].Name)
	//}

	//AIRWAYS
	//for i:= 0; i < len(airways); i++ {
	//	fmt.Println("\nName: " + airways[i].Name)
	//	fmt.Println(airways[i].Fixes)
	//}

	route := "KDFW HUDAD2 HUDAD PNH ZIGEE NIIXX2 KDEN"
}

func parseFIX(filename string) Fixes {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var fixes Fixes
	xml.Unmarshal(byteValue, &fixes)

	return fixes
}

func parseAWY(filename string) []Airway {
	txtFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer txtFile.Close()
	scanner := bufio.NewScanner(txtFile)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	var airways []Airway
	for _, eachline := range txtlines {
		name := eachline[1:strings.Index(eachline, "F")]
		fixes := strings.Fields(eachline[strings.Index(eachline, "F")+6:])

		obj := Airway{
			Name: name,
			Fixes: fixes,
		}
		airways = append(airways, obj)
	}

	return airways
}

func routedist(route string) uint {
	intersections := strings.Fields(route)
	dist := 0
	for i := 0; i < len(intersections)-1; i++ {

	}

	return dist
}