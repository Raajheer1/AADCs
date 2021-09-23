package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"github.com/umahmood/haversine"
	"io/ioutil"
	"os"
	"strings"
)

type Location struct {
	Lon float64 `xml:"Lon,attr"`
	Lat float64 `xml:"Lat,attr"`
}

type Fix struct {
	Name     string   `xml:"ID,attr"`
	Location Location `xml:"Location"`
}

type Fixes struct {
	Fixes []Fix `xml:"Waypoint"`
}

type Airway struct {
	Name  string
	Fixes []string
}

var fixes map[string]Location = parseFIX("Waypoints.xml")
var airways map[string][]string = parseAWY("AWY.txt")

/*
Callsign: ASA2118
Route: KSFO +SSTIK4 NTELL Q174 FLCHR COKTL1 KLAS
Distance: 13544.87
*/
//
//func main() {
//	//route := "KSFO +SSTIK4 NTELL Q174 FLCHR COKTL1 KLAS"
//	//route := "KSEA SEA7 SEA DCT NORMY J90 MWH/N0451F350 DCT KU87M DCT IDA DCT MJANE DCT KDEN"
//	//Need to filter out and remove the /N0451F350 before airways as it doesnt know when to end the airway
//	//fmt.Println(route)
//	//fmt.Println(Routeparse(route))
//}

func parseFIX(filename string) map[string]Location {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var fixesXML Fixes
	xml.Unmarshal(byteValue, &fixesXML)

	fixesmap := make(map[string]Location)

	for i := 0; i < len(fixesXML.Fixes); i++ {
		name := fixesXML.Fixes[i].Name
		loc := fixesXML.Fixes[i].Location
		fixesmap[name] = loc
	}

	return fixesmap
}

func parseAWY(filename string) map[string][]string {
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

	airwaysmap := make(map[string][]string)
	for _, eachline := range txtlines {
		name := eachline[1:strings.Index(eachline, "F")]
		fixeslist := strings.Fields(eachline[strings.Index(eachline, "F")+6:])

		airwaysmap[name] = fixeslist
	}

	return airwaysmap
}

func Routeparse(route string) []string {
	intersections := strings.Fields(route)
	for i := 0; i < len(intersections); i++ {
		if strings.Index(intersections[i], "/") > 0 {
			intersections[i] = intersections[i][0:strings.Index(intersections[i], "/")]
		}
	}
	var endroute []string

	//Need to add AWY conversion parser.
	for i := 0; i < len(intersections); i++ {
		if len(airways[intersections[i]]) != 0 {
			start := intersections[i-1]
			end := intersections[i+1]
			var airway []string
			between := false
			for _, s := range airways[intersections[i]] {
				if (s == start || s == end) && !between {
					between = true
					continue
				}
				if s == end {
					between = false
					break
				}
				if s == start {
					between = false
					for i := 0; i < len(airway)/2; i++ {
						j := len(airway) - i - 1
						airway[i], airway[j] = airway[j], airway[i]
					}
					break
				}
				if between {
					airway = append(airway, s)
				}
			}
			endroute = append(endroute, airway...)
		} else {
			endroute = append(endroute, intersections[i])
		}
	}

	for i, s := range endroute {
		if fixes[s].Lon == 0 {
			removeIndex(endroute, i)
		}
	}

	return endroute
}

func Routedist(route []string) float64 {
	dist := 0.0
	for i := 0; i < len(route)-1; i++ {
		point1 := haversine.Coord{Lat: fixes[route[i]].Lat, Lon: fixes[route[i]].Lon}
		point2 := haversine.Coord{Lat: fixes[route[i+1]].Lat, Lon: fixes[route[i+1]].Lon}
		mi, _ := haversine.Distance(point1, point2)
		mi /= 1.151
		dist += mi
	}

	return dist
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
