package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Choice",
		"GET",
		"/choices",
		ChoicesHandler,
	},
	Route{
		"Choice",
		"GET",
		"/choice",
		ChoiceHandler,
	},
	Route{
		"Play",
		"POST",
		"/play",
		PlayHandler,
	},
}
