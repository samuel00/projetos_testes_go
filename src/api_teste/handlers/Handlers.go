package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    setRequest(w);
    //json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
	    setRequest(w);
            return
        }
    }
    w.WriteHeader(http.StatusNotFound);    
    json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    setRequest(w);
    //json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        //json.NewEncoder(w).Encode(people)
    }
    setRequest(w); 
}


func setRequest(w http.ResponseWriter){	
    peopleJson , err := json.Marshal(people);
    if err != nil {
      panic(err);
    }
    w.Header().Set("Content-Type", "application/json");
    w.WriteHeader(http.StatusOK);
    w.Write(peopleJson);  
}

