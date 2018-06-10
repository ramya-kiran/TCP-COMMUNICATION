package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sroot)

	http.ListenAndServe(":9000", nil)

}

func sroot(response http.ResponseWriter, request *http.Request) {
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
