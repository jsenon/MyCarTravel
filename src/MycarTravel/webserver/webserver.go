package webserver

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Person struct {
	Name string //exported field since it begins with a capital letter
}

func Mytravelcarweb(res http.ResponseWriter, req *http.Request) {

	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>MyCarTravel</title>
		</head>
		<body>
		<form action="http://127.0.0.1:9000/mycartravel?maps=getdistance&thread=123" method="post" enctype="application/x-www-form-urlencoded">
		<input type="text" name="Origin" value="Toulouse"/>
		<input type="text" name="Destination" value="Nantes"/>
		<input type="submit"/>
		</form>
		</body>
		</html>`,
	)
	// The Form struct is a map, which keys are strings and values are a slice of strings if we use ParseForm
	// A struct that contains 2 maps if we use ParseMultipartForm The first map has keys that are strings and values that are slices of string while the second map is empty.
	// map[maps:[getdistance] thread:[123] Origin:[Toulouse] Destination:[Nantes]]
	req.ParseForm()
	//fmt.Fprintln(res, req.Form)
	fmt.Fprintln(res, req.PostForm)
	fmt.Fprintln(res, req.PostForm["Destination"])
	fmt.Fprintln(res, req.PostForm["Origin"])

}

func Getdistance(res http.ResponseWriter, req *http.Request) {
	t := template.New("custom template")
	t, _ = t.ParseFiles("templates/indexbis.html")
	user := Person{Name: "Julien"}
	t.Execute(res, user)
}

func Index(res http.ResponseWriter, req *http.Request) {
	Render(res, "templates/index.html", nil)
}

func Render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
