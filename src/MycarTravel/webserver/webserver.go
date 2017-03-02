package webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

/*// Function for Rendering templates
func Render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

var MyApiKey = os.Getenv("GOOGLE_APIKEY")

// Starting point of MyTravelCar
func Index(res http.ResponseWriter, req *http.Request) {
	const tpl = `
	<!DOCTYPE html>
	 	<html>
		<head>
	  		<meta charset="UTF-8">
	  		<title>{{.Title}}</title>
		</head>
	    <body>
	    	<h1>{{.Title}}</h1>
	    	For information you have set Google API key to: {{.Myapikey}}
	    	</br>
	    	</br>
	    	<form action="/send" method="POST" novalidate>
	    		<div>
	    			<label>Your Origin:</label>
	    			<input type="text" name="origin">
	    		</div>
	    		<div>
	    			<label>Your Destination:</label>
	    			<input type="text" name="destination">  
	    		</div>
	    		<div>
	    			<input type="submit" value="Send Calculation">
	    		</div>
	    	</form>
	   </body>
	 </html>`
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)
	data := struct {
		Title    string
		Myapikey string
	}{
		Title:    "MyCarTravel",
		Myapikey: "toto",
	}
	data.Myapikey = MyApiKey
	err = t.Execute(res, data)
	check(err)
}

// For future Usage, Check and validate fields and redirect to results page
func CheckFields(res http.ResponseWriter, req *http.Request) {
	// Retrieve Input User from Index func
	// Launch API Google with input
	// Redirect to results with google's answer
	var Depart string
	var Finish string
	var mapsanswer DistanceMatrixResponse
	req.ParseForm()
	//fmt.Fprintln(res, req.Form)
	//fmt.Fprintln(res, req.FormValue("origin"))
	Depart = req.FormValue("origin")
	Finish = req.FormValue("destination")
	url, erro := http.Get("https://maps.googleapis.com/maps/api/distancematrix/json" + "?units=metric&origins=" + Depart + "&destinations=" + Finish + "&key=" + MyApiKey)
	if erro != nil {
		// handle error
		fmt.Println("error:", erro)
		fmt.Println("res:", res)
		return
	}
	defer url.Body.Close()
	if erro := json.NewDecoder(url.Body).Decode(&mapsanswer); erro != nil {
		return
	}
	//fmt.Fprintln(res, url)
	fmt.Fprintln(res, mapsanswer.OriginAddresses[0])
	fmt.Fprintln(res, mapsanswer.DestinationAddresses[0])
	fmt.Fprintln(res, mapsanswer.Rows[0].Elements[0].Duration.Text)
	fmt.Fprintln(res, mapsanswer.Rows[0].Elements[0].Distance.HumanReadable)
	//http.Redirect(res, req, "/results", http.StatusSeeOther)
}

// Results Page
func Results(res http.ResponseWriter, req *http.Request) {
	const tpl = `
	<!DOCTYPE html>
	 	<html>
		<head>
	  		<meta charset="UTF-8">
	  		<title>{{.Title}}</title>
		</head>
	    <body>
	    	<h1>{{.Title}}</h1>
	   </body>
	 </html>`
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)
	data := struct {
		Title string
	}{
		Title: "MyCarTravel Results",
	}
	err = t.Execute(res, data)
	check(err)
}
