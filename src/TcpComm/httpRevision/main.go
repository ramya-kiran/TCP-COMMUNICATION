package main

import (
	"io"
	"net/http"
	"strings"
)

type dogHandler int

type catHandler int

/*
The main function initializes a multiplexer, which routes
/dog/ patter URL to dog handler and
/cat/ patter URL to cat handler
*/
func main() {

	var dog dogHandler
	var cat catHandler

	mux := http.NewServeMux()

	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)

	// The server listens and serves all requests coming to to locahost:9000
	// using mux as the multiplexer
	http.ListenAndServe(":9000", mux)

}

func (d dogHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	fs := strings.Split(request.URL.Path, "/")
	if len(fs) >= 3 {
		dogname := fs[2]
		io.WriteString(response, `<strong> Welcome to this webpage </strong> with dog name as `+dogname+`</br>
			<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
			`)
	} else {

		io.WriteString(response, `<strong> Welcome to this webpage </strong>`)
	}
}

func (c catHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	fs := strings.Split(request.URL.Path, "/")
	if len(fs) >= 3 {
		dogname := fs[2]
		io.WriteString(response, `<strong> Welcome to this webpage </strong> with dog name as `+dogname+`</br>
			<img src="https://upload.wikimedia.org/wikipedia/commons/0/06/Kitten_in_Rizal_Park%2C_Manila.jpg">
			`)
	} else {

		io.WriteString(response, `<strong> Welcome to this webpage through cat handler </strong>`)
	}
}
