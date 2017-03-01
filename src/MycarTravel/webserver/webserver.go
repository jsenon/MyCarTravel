package webserver

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
