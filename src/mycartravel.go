//https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=Washington,DC&destinations=New+York+City,NY&key=YOURAPIKEY

package main

import (
	"fmt"
)

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

// Define a struct for api distance matrix answer
type TripAnswer struct {
	Destination string
	Origin      string
	answer      map[string]map[string]map[string]map[string]float64
	status      string
}

/*
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
	env := []Town{
		{"Toulouse", "Toulouse"},
	}
	fmt.Println("Enter Origin:")
	fmt.Scanf("%s", &env[0].TownFrom)
	fmt.Println("Enter Destination:")
	fmt.Scanf("%s", &env[0].TownTo)
	fmt.Println("Trip", env)
}
