package main

import "time"

type Todo struct {
	Id int			`json:"id"`
	Name string 	`json:"name"`
	Completed bool	`json:"completed"`
	Due time.Time	`json:"due"`
}

type Updatetodo struct {
	Feilds []string	`json: "feilds"`
	Data	Todo	`json: "data"`
}

type Todos []Todo