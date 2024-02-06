package geospatial

import (
	"fmt"
	"testing"
)

var privatekey = ""
var publickey = ""
var encode = ""
var dbname = "pointmap"
var collname = "pointmap"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{106.82320902216969, -6.180906971485783},
				{106.82333238308496, -6.183145216209596},
				{106.83304705518373, -6.1825013385601295},
				{106.83212184831717, -6.180447057013481},
				{106.82320902216969, -6.180906971485783},
			},
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{95.31123456789012, 5.553210987654321},
				{95.31133456789011, 5.553210987654321},
				{95.31133456789011, 5.553310987654321},
				{95.31123456789012, 5.553310987654321},
				{95.31123456789012, 5.553210987654321},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "geojson", "bandaaceh")
	coordinates := Point{
		Coordinates: []float64{
			95.30987654321098, 5.556789012345678,
		},
	}
	datagedung := Near(mconn, "bandaaceh", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "geojson")
	coordinates := Point{
		Coordinates: []float64{
			95.30987654321098, 5.556789012345678,
		},
	}
	datagedung := NearSphere(mconn, "bandaaceh", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "geojson")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{95.32345678901234, 5.567890123456789},
			{95.32355678901234, 5.567990123456789},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}