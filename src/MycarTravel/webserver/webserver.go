package webserver

import (
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
	    			<input type="text" name="origin">  
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
	// Future Check fields
	// Launch API Google
	http.Redirect(res, req, "/results", http.StatusSeeOther)
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
