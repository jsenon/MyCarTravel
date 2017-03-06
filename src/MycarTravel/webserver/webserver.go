package webserver

import (
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Function for Rendering templates
// filename is relative path form where you run the bin

func Render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

var MyApiKey = os.Getenv("GOOGLE_APIKEY")
var Mapsanswer DistanceMatrixResponse

// Starting point of MyTravelCar
func Index(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title    string
		Myapikey string
	}{
		Title:    "MyCarTravel",
		Myapikey: "toto",
	}
	data.Myapikey = MyApiKey
	Render(res, "src/templates/index.html", data)
}

// For future Usage, Check and validate fields and redirect to results page
func CheckFields(res http.ResponseWriter, req *http.Request) {
	// Retrieve Input User from Index func
	// Launch API Google with input
	// Redirect to results with google's answer
	var Depart string
	var Finish string
	var Travelmode string

	req.ParseForm()
	Depart = req.FormValue("origin")
	Finish = req.FormValue("destination")
	Travelmode = req.FormValue("travelmode")

	fmt.Println("Depart: ", Depart)
	fmt.Println("Finish: ", Finish)

	// Check if Origin and Destination is not blank
	if len(req.Form["origin"][0]) == 0 || len(req.Form["destination"][0]) == 0 {
		http.Redirect(res, req, "/error", http.StatusSeeOther)
		fmt.Println("ORIGIN OR DESTINATION NULL")
		return
	}

	Depart2 := strings.Replace(Depart, " ", "-", -1)
	fmt.Println("Depart Replace: ", Depart2)

	Finish2 := strings.Replace(Finish, " ", "-", -1)
	fmt.Println("Depart Replace: ", Finish2)

	// Input English Letters
	if m, _ := regexp.MatchString("^[a-zA-Z-']+$", Depart2); !m {
		http.Redirect(res, req, "/error", http.StatusSeeOther)
		fmt.Println("NON ENGLISH CHARACTERS")
		return
	}
	if m, _ := regexp.MatchString("^[a-zA-Z-']+$", Finish2); !m {
		http.Redirect(res, req, "/error", http.StatusSeeOther)
		fmt.Println("NON ENGLISH CHARACTERS")
		return
	}

	//Escape particular strings. It escapes only five such characters: <, >, &, '
	fmt.Println(html.EscapeString(Finish2))

	url, erro := http.Get("https://maps.googleapis.com/maps/api/distancematrix/json" + "?units=metric&origins=" + Depart2 + "&destinations=" + Finish2 + "&mode=" + Travelmode + "&key=" + MyApiKey)

	if erro != nil {
		// handle error
		http.Redirect(res, req, "/error", http.StatusSeeOther)
		fmt.Println("error:", erro)
		fmt.Println("res:", res)
		return
	}
	defer url.Body.Close()
	if erro := json.NewDecoder(url.Body).Decode(&Mapsanswer); erro != nil {
		return
	}

	fmt.Println("Map elements", Mapsanswer.Rows[0].Elements[0].Status)

	// Case OK NOT_FOUND ZERO_RESULTS
	switch Mapsanswer.Rows[0].Elements[0].Status {
	case "OK":
		http.Redirect(res, req, "/results", http.StatusSeeOther)
	case "NOT_FOUND":
		http.Redirect(res, req, "/notcovered", http.StatusSeeOther)
	case "ZERO_RESULTS":
		http.Redirect(res, req, "/notcovered", http.StatusSeeOther)
	default:
		http.Redirect(res, req, "/error", http.StatusSeeOther)
	}

}

// Results Page
func Results(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title       string
		Origin      string
		Destination string
		Duration    string
		Distance    string
	}{
		Title:       "MyCarTravel Results",
		Origin:      Mapsanswer.OriginAddresses[0],
		Destination: Mapsanswer.DestinationAddresses[0],
		Duration:    Mapsanswer.Rows[0].Elements[0].Duration.Text,
		Distance:    Mapsanswer.Rows[0].Elements[0].Distance.HumanReadable,
	}

	// Use Mapsanswer for results screen
	Render(res, "src/templates/results.html", data)
}

func Htmltemplate(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title string
	}{
		Title: "MyCarTravel Results",
	}
	Render(res, "src/templates/test.html", data)
}

func Wrong(res http.ResponseWriter, req *http.Request) {
	Render(res, "src/templates/error.html", nil)
}

func Error(res http.ResponseWriter, req *http.Request) {
	Render(res, "src/templates/404.html", nil)
}

func NotCovered(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title string
		Error string
	}{
		Title: "MyCarTravel Results",
		Error: "Sorry, your search appears to be outside our current coverage area for driving.",
	}
	Render(res, "src/templates/notcovered.html", data)
}
