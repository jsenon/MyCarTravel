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

// FIX ME
// Error with Dest and ori with space

/*
https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=Toulouse&destinations=Nantes&key=YOURAPIKEY
*/

func main() {
	// Redirect Page
	http.HandleFunc("/hello", webserver.Index)
	http.HandleFunc("/", webserver.Error)
	http.HandleFunc("/send", webserver.CheckFields)
	http.HandleFunc("/results", webserver.Results)
	http.HandleFunc("/error", webserver.Wrong)
	// Test Web page from template
	http.HandleFunc("/htmltemplate", webserver.Htmltemplate)
	// Init WebServer
	http.ListenAndServe(":9000", nil)

}
