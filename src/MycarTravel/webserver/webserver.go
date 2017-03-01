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
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>MyCarTravel</title>
		</head>
		<body>
		<form action="http://127.0.0.1:9000/mycartravel?hello=world&thread=123" method="post" enctype="application/x-www-form-urlencoded">
		<input type="text" name="Origin" value="Toulouse"/>
		<input type="text" name="Destination" value="Nantes"/>
		<input type="submit"/>
		</form>
		</body>
		</html>`,
	)
}
