package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

var people []Person

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    peopleJson , err := json.Marshal(persons);
    if err != nil {
         	panic(err);
    }
    setRequest(w,peopleJson);
    //json.NewEncoder(w).Encode(persons)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    person := RepoFindPerson(params["id"])
    if person.ID != "" {
	 peopleJson , err := json.Marshal(person);
         if err != nil {
         	panic(err);
    	}
        setRequest(w,peopleJson)    	
        return;
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
    persons = append(persons, person)
    peopleJson , err := json.Marshal(persons);
    if err != nil {
     	panic(err);
    }
    setRequest(w,peopleJson)
    //json.NewEncoder(w).Encode(persons)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    RepoDestroyPerson(params["id"])
    
    peopleJson , err := json.Marshal(persons);
    if err != nil {
         	panic(err);
    }
    setRequest(w,peopleJson) 
}

func setRequest(w http.ResponseWriter, peopleJson []uint8){	
    w.Header().Set("Content-Type", "application/json");
    w.WriteHeader(http.StatusOK);
    w.Write(peopleJson);  
}

