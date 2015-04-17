package hateoas

import (
	"net/http"
	"testing"
)

// Test resource
type testResource struct {
	Id string
}

func (r testResource) MarshalJSON() ([]byte, error) {
	return nil, nil
}
func (r testResource) UnmarshalJSON(b []byte) error {
	return nil
}

// Test resource handler
type testResourceHandler struct {
}

func (rh testResourceHandler) ResourceName() string {
	return "resources"
}
func (rh testResourceHandler) GetOne(id string) (Resource, *Error) {
	testResource := testResource{}
	return testResource, nil
}
func (rh testResourceHandler) GetAll(pageOpts PageOpts) ([]Resource, *Error) {
	testResources := []testResource{}

	resources := make([]Resource, len(testResources))
	for index, value := range testResources {
		resources[index] = value
	}

	return resources, nil
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
