package webserver

import (
	"io"
	"net/http"
)

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
		    <title>MyCarTravel</title>
		  </head>
		  <body>
		    MycarTravel
		  </body>
		</html>`,
	)
}
