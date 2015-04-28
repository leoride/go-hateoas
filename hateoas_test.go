package hateoas

import (
	"net/http"
	"testing"
)

// Test resource
type testResourceAnimal struct {
	Id   string
	Name string
}

func (tr testResourceAnimal) GetId() string {
	return tr.Id
}

type testResourcePerson struct {
	Id     string
	Name   string
	Animal testResourceAnimal
}

func (tr testResourcePerson) GetId() string {
	return tr.Id
}

// Test resource handler
type testResourcePersonHandler struct {
}

func (rh testResourcePersonHandler) ResourceName() string {
	return "resources"
}
func (rh testResourcePersonHandler) Count() (int, *Error) {
	return 25, nil
}

func (rh testResourcePersonHandler) GetOne(id string) (Resource, *Error) {
	testResourcePerson := testResourcePerson{}
	return testResourcePerson, nil
}
func (rh testResourcePersonHandler) GetAll(pageOpts PageOpts) ([]Resource, *Error) {
	testResourcePersons := make([]testResourcePerson, 5)

	smoky := testResourceAnimal{"1", "smoky"}

	testResourcePersons[0] = testResourcePerson{"1", "tomg", smoky}
	testResourcePersons[1] = testResourcePerson{"2", "jacquesg", testResourceAnimal{}}
	testResourcePersons[2] = testResourcePerson{"3", "isabelleg", testResourceAnimal{}}
	testResourcePersons[3] = testResourcePerson{"4", "marieg", testResourceAnimal{}}
	testResourcePersons[4] = testResourcePerson{"5", "lylwenng", testResourceAnimal{}}

	resources := make([]Resource, len(testResourcePersons))

	i := 0
	for index, value := range testResourcePersons {
		if (pageOpts.Offset <= index) && (index < (pageOpts.Offset + pageOpts.Limit)) {
			resources[i] = value
			i++
		}
	}

	return resources[0:i], nil
}
func (rh testResourcePersonHandler) Create(newR Resource) (string, *Error) {
	return "", nil
}
func (rh testResourcePersonHandler) Update(id string, updR Resource) (Resource, *Error) {
	testResourcePerson := updR.(testResourcePerson)
	return testResourcePerson, nil
}
func (rh testResourcePersonHandler) Delete(string) *Error {
	return nil
}

// This test runs: go to http://localhost:8080/api/resources to test
func TestHandle(t *testing.T) {
	testResourcePersonHandler := testResourcePersonHandler{}
	Handle("/api", testResourcePersonHandler)
	http.ListenAndServe(":8080", nil)
}
