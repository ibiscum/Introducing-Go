package main

import "net/http"

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	res.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	//Handle the /hello route by calling the hello function
	http.HandleFunc("/hello", hello)

	//Serve on port 9000
	http.ListenAndServe(":9000", nil)

	//Handle static files by using FileServer
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}
