//
// Based on Google Maps API
//

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// TO DO
// Construct API Call with parameters
// Build Web Server
// Build ouside main a type package

// Define Price
type GazPrice struct {
	TypeGas string
	Price   float64
	Devise  string
}

// Answer Structure coming from google API
type DistanceMatrixResponse struct {
	// OriginAddresses contains an array of addresses as returned by the API from your original request.
	OriginAddresses []string `json:"origin_addresses"`
	// DestinationAddresses contains an array of addresses as returned by the API from your original request.
	DestinationAddresses []string `json:"destination_addresses"`
	// Rows contains an array of elements.
	Rows []DistanceMatrixElementsRow `json:"rows"`
}

// DistanceMatrixElementsRow is a row of distance elements.
type DistanceMatrixElementsRow struct {
	//Elements contains an array Distance Duration and Status
	Elements []*DistanceMatrixElement `json:"elements"`
}

type DistanceMatrixElement struct {
	// Elements Status Code
	Status string `json:"status"`
	// Duration is the length of time it takes to travel this route.
	Duration Duration `json:"duration"`
	// DurationInTraffic is the length of time it takes to travel this route considering traffic.
	DurationInTraffic Duration `json:"duration_in_traffic"`
	// Distance is the total distance of this route.
	Distance Distance `json:"distance"`
}

type Distance struct {
	// HumanReadable is the human friendly distance. This is rounded and in an appropriate unit for the
	// request. The units can be overriden with a request parameter.
	HumanReadable string `json:"text"`
	// Meters is the numeric distance, always in meters. This is intended to be used only in
	// algorithmic situations, e.g. sorting results by some user specified metric.
	Meters int `json:"value"`
}

type Duration struct {
	// Value indicates the duration, in seconds.
	Value int64 `json:"value"`
	// Text contains a human-readable representation of the duration.
	Text string `json:"text"`
}

// Mode is for specifying travel mode. Default dirving
type Mode string

// Avoid is for specifying routes that avoid certain features. Values tolls, highways,ferries,indoor
type Avoid string

// Units specifies which units system to return human readable results in. Metric or Imperial
type Units string

// TransitMode is for specifying a transit mode for a request
type TransitMode string

// TrafficModel specifies traffic prediction model when requesting future directions.
type TrafficModel string

// Travel mode preferences.
const (
	TravelModeDriving   = Mode("driving")
	TravelModeWalking   = Mode("walking")
	TravelModeBicycling = Mode("bicycling")
	TravelModeTransit   = Mode("transit")
)

// Features to avoid.
const (
	AvoidTolls    = Avoid("tolls")
	AvoidHighways = Avoid("highways")
	AvoidFerries  = Avoid("ferries")
)

// Units to use on human readable distances.
const (
	UnitsMetric   = Units("metric")
	UnitsImperial = Units("imperial")
)

// Transit mode of directions or distance matrix request.
const (
	TransitModeBus    = TransitMode("bus")
	TransitModeSubway = TransitMode("subway")
	TransitModeTrain  = TransitMode("train")
	TransitModeTram   = TransitMode("tram")
	TransitModeRail   = TransitMode("rail")
)

// Traffic prediction model when requesting future directions.
const (
	TrafficModelBestGuess   = TrafficModel("best_guess")
	TrafficModelOptimistic  = TrafficModel("optimistic")
	TrafficModelPessimistic = TrafficModel("pessimistic")
)

/*
https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=Toulouse&destinations=Nantes&key=YOURAPIKEY

units=metric (default) returns distances in kilometers and meters.
units=imperial returns distances in miles and feet.

{
   "destination_addresses" : [ "Nantes, France" ],
   "origin_addresses" : [ "Toulouse, France" ],
   "rows" : [
      {
         "elements" : [
            {
               "distance" : {
                  "text" : "363 mi",
                  "value" : 584955
               },
               "duration" : {
                  "text" : "5 hours 14 mins",
                  "value" : 18846
               },
               "status" : "OK"
            }
         ]
      }
   ],
   "status" : "OK"
}

*/

func main() {

	// Initialization
	var mapsanswer DistanceMatrixResponse
	var Addfrom string
	var Addto string

	// Export your API Key
	apiKey := os.Getenv("GOOGLE_APIKEY")
	fmt.Println("your key", apiKey)
	// Ask From
	fmt.Println("Enter your Origin Address: ")
	fmt.Scanf("%s", &Addfrom)
	// Ask To
	fmt.Println("Enter your Destination Address: ")
	fmt.Scanf("%s", &Addto)
	// Launch connexion to maps google
	res, erro := http.Get("https://maps.googleapis.com/maps/api/distancematrix/json" + "?units=metric&origins=" + Addfrom + "&destinations=" + Addto + "&key=" + apiKey)
	if erro != nil {
		// handle error
		return
	}
	defer res.Body.Close()
	if erro := json.NewDecoder(res.Body).Decode(&mapsanswer); erro != nil {
		return
	}
	fmt.Println(mapsanswer.OriginAddresses[0])
	fmt.Println(mapsanswer.DestinationAddresses[0])
	fmt.Println(mapsanswer.Rows[0].Elements[0].Duration.Text)
	fmt.Println(mapsanswer.Rows[0].Elements[0].Distance.HumanReadable)

}
