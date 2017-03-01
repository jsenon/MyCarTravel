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
		    <title>Hello, World</title>
		  </head>
		  <body>
		    Hello, world!
		  </body>
		</html>`,
	)
}

//func main() {
//	http.HandleFunc("/hello", hello)
//	http.ListenAndServe(":9000", nil)
//}