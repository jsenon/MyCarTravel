//https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=Washington,DC&destinations=New+York+City,NY&key=YOURAPIKEY

package main

// Define a struct for api distance matrix
type Trip struct {
	TownFrom string
	TownTo   string
}

// Define a struct for api distance matrix answer
type TripAnswer struct {
	Destination string
	Origin string
	answer map[]string
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
	var trip
	printf
	scanf
}
