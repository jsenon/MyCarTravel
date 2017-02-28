package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
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

// Define My Query
/*var distanceMatrixAPI = &apiConfig{
	host:            "https://maps.googleapis.com",
	path:            "/maps/api/distancematrix/json",
	acceptsClientID: true,
}*/

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
	Elements []*DistanceMatrixElement `json:"elements"`
}

// Definition of DistanceElement
// Define time.Duration and uncomment
// Define distance and uncomment
type DistanceMatrixElement struct {
	Status string `json:"status"`
	// Duration is the length of time it takes to travel this route.
	//Duration time.Duration `json:"duration"`
	// DurationInTraffic is the length of time it takes to travel this route considering traffic.
	//DurationInTraffic time.Duration `json:"duration_in_traffic"`
	// Distance is the total distance of this route.
	//Distance Distance `json:"distance"`
}

/*
https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=Toulouse&destinations=Nantes&key=YOURAPIKEY

{
   "destination_addresses" : [ "New York, NY, USA" ],
   "origin_addresses" : [ "Washington, DC, USA" ],
   "rows" : [
      {
         "elements" : [
            {
               "distance" : {
                  "text" : "225 mi",
                  "value" : 361722
               },
               "duration" : {
                  "text" : "3 hours 48 mins",
                  "value" : 13672
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

	// Initialization?
	var mapsanswer DistanceMatrixResponse

	// Export your API Key
	apiKey := os.Getenv("GOOGLE_APIKEY")

	fmt.Println("your key", apiKey)
	// Launch connexion to maps google
	res, erro := http.Get("https://maps.googleapis.com/maps/api/distancematrix/json" + "?units=imperial&origins=Toulouse&destinations=Nantes" + "&key=" + apiKey)
	defer res.Body.Close()
	if erro != nil {
		panic(erro.Error())
	}
	err := json.NewDecoder(res.Body).Decode(&mapsanswer)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(mapsanswer.OriginAddresses)
	fmt.Println(mapsanswer.DestinationAddresses)
}
