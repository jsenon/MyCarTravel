//
// Based on Google Maps API
//

package main

import (
	"MycarTravel/webserver"
	"net/http"
)

// TO DO
// result in HTML
// use template instead inline HTML
// catch error exeption
// ask option values
// sublime HTML page

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
	// Redirect Page
	http.HandleFunc("/hello", webserver.Index)
	http.HandleFunc("/send", webserver.CheckFields)
	http.HandleFunc("/results", webserver.Results)
	http.HandleFunc("/error", webserver.Wrong)
	// Test Web page from template
	http.HandleFunc("/htmltemplate", webserver.Htmltemplate)

	// Init WebServer
	http.ListenAndServe(":9000", nil)

}
