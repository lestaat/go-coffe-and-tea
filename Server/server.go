package main

import ("net/http")

func main () {

	// Route
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/test", homeHandlerTest)

	// Server
	http.ListenAndServe(":3000", nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello World"))
}

func homeHandlerTest(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Aca"))
}