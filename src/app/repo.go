package main

import (
	"fmt"
)

var todos Todos
var currentId int 

func init(){
	RepoCreateTodo(Todo{Name: "Write presention"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

// create

func RepoCreateTodo(todo Todo) Todo {
	currentId += 1
	todo.Id = currentId
	todos = append(todos,todo)

	return todo
}

// read

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}

	return Todo{}
}

// update

func RepoUpdateTodo(id int , data Updatetodo) Todo {
	
	feilds := data.Feilds
	todo := data.Data

	fmt.Println(feilds,todo)

	for i , t := range todos {
		if t.Id == id {
			
			for _,feild := range feilds {
				
				switch feild {
				case "name":
					todos[i].Name = todo.Name
					fmt.Println("name",todos[i])
				
				case "completed" :
					todos[i].Completed = todo.Completed
					fmt.Println("completed",todos[i])

				case "due" :
					todos[i].Due = todo.Due
					fmt.Println("due",todos[i])					
				}
			}

			return todos[i]
		}
	}

	return Todo{}
}

// delete

func RepoDestroyTodo(id int) (Todo , error) {

	for i,t := range todos {
		if t.Id == id {
			todos = append(todos[:i],todos[i+1:]...)
			return t,nil
		}
	}

	return Todo{},fmt.Errorf("id not found")
}