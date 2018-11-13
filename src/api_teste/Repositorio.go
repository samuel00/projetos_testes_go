package main

import ("fmt"
	"strconv")

var currentId int

var persons []Person
var person []Person

// Give us some seed data
func init() {

    RepoCreatePerson(Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    RepoCreatePerson(Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
}

func RepoFindPerson(id string) Person {
    for _, p := range persons {
        if p.ID == id {
            return p
        }
    }
    // return empty Todo if not found
    return Person{}
}

func RepoCreatePerson(p Person) Person {
    currentId += 1
    s := strconv.Itoa(currentId)
    p.ID = s
    persons = append(persons, p)
    return p
}

func RepoDestroyPerson(id string) error {
    for i, p := range persons {
        if p.ID == id {
            persons = append(persons[:i], persons[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
