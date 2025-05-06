package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiconfig) handlerPass(w http.ResponseWriter, req *http.Request) {
	// Set the Content-Type header to serve HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the HTML content
	html := `
	<!DOCTYPE html>
	<html>
	<head>
			<title>My Web Page</title>
	</head>
	<body>
			<h1>Welcome to My Web Page</h1>
			<p> Hi Docker, I pushed a new version! </p>
	</body>
	</html>
	`

	fmt.Fprint(w, html)
}