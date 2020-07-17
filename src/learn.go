package main

import (
	"fmt"
	_ "reflect"
)

type Demo struct {
	Id int
	Name string
}

type Demos []Demo

var one string = "1"

var demos = Demos{

	Demo{
		int(one),
		"hellp",
	},
	Demo{
		2,
		"henlo",
	},
	Demo{
		3,
		"hepno",
	},
	Demo{
		4,
		"penlo",
	},
}

func main()  {

	for _,d := range demos {
		fmt.Println(d.Id , d.Name)
	}

	for i := range demos {
		demos[i]."Name"= "mayank"
	}

	for _,d := range demos {
		fmt.Println(d.Id , d.Name)
	}
}
