package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"strconv"

	"github.com/gorilla/mux"
)

// /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Golang basic JSON CRUD API ")
}

//  /todos
func TodoIndex(w http.ResponseWriter , r *http.Request) {
	
	w.Header().Set("Content-Type","Application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

//  /todos/{todoId}
func TodoShow(w http.ResponseWriter , r *http.Request) {
    vars := mux.Vars(r)
	todoid := vars["todoId"]

	id,err := strconv.Atoi(todoid) 
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invlid Id")
	}else{
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(RepoFindTodo(id)); err != nil {
			panic(err)
		}
	}
	
}

func TodoDelete( w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	todoid := vars["todoId"]

	id,err := strconv.Atoi(todoid)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity) //unporcessable entity
		fmt.Fprintf(w,"Invalid id ")

	}else{

		t,err := RepoDestroyTodo(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //unporcessable entity
			fmt.Fprintf(w,"Id doesn't Exist")
		}else{
			w.Header().Set("Content-Type","application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK) //unporcessable entity

			if err := json.NewEncoder(w).Encode(t); err != nil {
				panic(err)
			}
		}
	}
}

func TodoCreate(w http.ResponseWriter , r *http.Request) {

	var todo Todo 
	body , err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil{
		panic(err)
	}


	if err := json.Unmarshal(body, &todo); err != nil{
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		w.WriteHeader(422) //unporcessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type","appplication/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func TodoUpdate(w http.ResponseWriter , r *http.Request) {

	vars := mux.Vars(r)
	todoid := vars["todoId"]

	id,err := strconv.Atoi(todoid)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity) //unporcessable entity
		fmt.Fprintf(w,"Invalid id ")

	}else{

		var data Updatetodo
		body , err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))

		if err != nil {
			panic(err)
		}

		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(body,&data); err != nil {
			w.Header().Set("Content-Type","application/json; charset=UTF-8")
			w.WriteHeader(422) //unporcessable entity

			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		fmt.Println(data)

		t := RepoUpdateTodo(id,data)
		w.Header().Set("Content-Type","appplication/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	}
}