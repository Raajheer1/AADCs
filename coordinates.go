package main

import "math"

type Coordiante struct {
	Lon float64
	Lat float64
}

type CartesianCoordinate struct {
	X float64
	Y float64
	Z float64
}

type Line struct {
	start CartesianCoordinate
	end   CartesianCoordinate
}

var R float64 = 6371

func ConvertCoordinate(coord Coordiante) CartesianCoordinate {
	x := R * math.Cos(coord.Lat*math.Pi/180.0) * math.Cos(coord.Lon*math.Pi/180.0)
	y := R * math.Cos(coord.Lat*math.Pi/180.0) * math.Sin(coord.Lon*math.Pi/180.0)
	z := R * math.Sin(coord.Lat*math.Pi/180.0)
	return CartesianCoordinate{
		X: x,
		Y: y,
		Z: z,
	}
}

func ConvertCartesian(cart CartesianCoordinate) Coordiante {
	lat := math.Asin(cart.Z / R)
	lon := math.Atan2(cart.Y, cart.X)
	return Coordiante{
		Lat: lat,
		Lon: lon,
	}
}

func getPoints(line Line, num int) []CartesianCoordinate {
	var points [num]CartesianCoordinate
	ydiff := line.end.Y - line.start.Y
	xdiff := line.end.X - line.start.X
	slope := line.end.Y -
	return points
}
