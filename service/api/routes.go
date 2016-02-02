package api

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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Classify",
		"GET",
		"/classify",
		Classify,
	},
	Route{
		"Analyze",
		"GET",
		"/analyze",
		Analyze,
	},
	Route{
		"ShowTasks",
		"GET",
		"/show/tasks",
		ShowTasks,
	},
}