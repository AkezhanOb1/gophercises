package main

import (

	"net/http"
	"Go-Web/go-exercise/storyJsonParser/router"
)




func main() {

	http.HandleFunc("/", router.Intro)
	http.HandleFunc("/new-york", router.NewYork)
	http.HandleFunc("/debate", router.Debate)
	http.HandleFunc("/seanKelly", router.SeanKelly)
	http.HandleFunc("/mark-bates", router.MarkBates)
	http.HandleFunc("/denver", router.Denver)
	http.HandleFunc("/home", router.Home)


	http.ListenAndServe(":8080", nil)
}



