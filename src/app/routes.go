package main

import (
	"net/http"
)

type Route struct {
	Name string
	Method string
	Path string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes

func init(){

	routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},
		Route{
			"TodoIndex",
			"GET",
			"/todos",
			TodoIndex,
		},
		Route{
			"TodoShow",
			"GET",
			"/todo/{todoId}",
			TodoShow,
		},
		Route{
			"TodoDelete",
			"DELETE",
			"/todo/{todoId}",
			TodoDelete,
		},			
		Route{
			"TodoCreate",
			"PUT",
			"/todo",
			TodoCreate,
		},
		Route{
			"TodoUpdate",
			"POST",
			"/todo/{todoId}",
			TodoUpdate,
		},
	}
}