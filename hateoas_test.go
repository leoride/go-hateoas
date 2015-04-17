package hateoas

import (
	"net/http"
	"testing"
)

type testResourceHandler struct {
}

func (rh testResourceHandler) ResourceName() string {
	return "resources"
}
func (rh testResourceHandler) GetOne(string) (Resource, *Error) {
	return nil, nil
}
func (rh testResourceHandler) GetAll(PageOpts) ([]Resource, *Error) {
	return nil, nil
}
func (rh testResourceHandler) Create(Resource) (string, *Error) {
	return "", nil
}
func (rh testResourceHandler) Update(string, Resource) (Resource, *Error) {
	return nil, nil
}
func (rh testResourceHandler) Delete(string) *Error {
	return nil
}

func TestHandle(t *testing.T) {
	testResourceHandler := testResourceHandler{}
	Handle("/api", testResourceHandler)
	http.ListenAndServe(":8080", nil)
}
