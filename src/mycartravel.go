package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// TO DO
// How to ask google map
// How to retrieve toll
// Build Web Server
// Fisrt Version with API Call
// Seconf Version with google map go plugin
// Define a struct for api distance matrix

type Town struct {
	TownFrom string
	TownTo   string
}

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
	// Export your API Key
	apiKey := os.Getenv("GOOGLE_APIKEY")
	fmt.Println("your key", apiKey)
	// Launch connexion to maps google
	res, erro := http.Get("https://maps.googleapis.com/maps/api/distancematrix/json" + "?units=metric&origins=Toulouse&destinations=Nantes" + "&key=" + apiKey)
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
