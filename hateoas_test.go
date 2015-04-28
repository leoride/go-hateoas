package hateoas

import (
	"net/http"
	"testing"
)

// Test resource
type testResource struct {
	Id   string
	Name string
}

func (tr testResource) GetId() string {
	return tr.Id
}

// Test resource handler
type testResourceHandler struct {
}

func (rh testResourceHandler) ResourceName() string {
	return "resources"
}
func (rh testResourceHandler) Count() (int, *Error) {
	return 25, nil
}

func (rh testResourceHandler) GetOne(id string) (Resource, *Error) {
	testResource := testResource{}
	return testResource, nil
}
func (rh testResourceHandler) GetAll(pageOpts PageOpts) ([]Resource, *Error) {
	testResources := make([]testResource, 5)
	testResources[0] = testResource{"1", "tomg"}
	testResources[1] = testResource{"2", "jacquesg"}
	testResources[2] = testResource{"3", "isabelleg"}
	testResources[3] = testResource{"4", "marieg"}
	testResources[4] = testResource{"5", "lylwenng"}

	resources := make([]Resource, len(testResources))

	i := 0
	for index, value := range testResources {
		if (pageOpts.Offset <= index) && (index < (pageOpts.Offset + pageOpts.Limit)) {
			resources[i] = value
			i++
		}
	}

	return resources[0:i], nil
}
func (rh testResourceHandler) Create(newR Resource) (string, *Error) {
	return "", nil
}
func (rh testResourceHandler) Update(id string, updR Resource) (Resource, *Error) {
	testResource := updR.(testResource)
	return testResource, nil
}
func (rh testResourceHandler) Delete(string) *Error {
	return nil
}

// This test runs: go to http://localhost:8080/api/resources to test
func TestHandle(t *testing.T) {
	testResourceHandler := testResourceHandler{}
	Handle("/api", testResourceHandler)
	http.ListenAndServe(":8080", nil)
}
