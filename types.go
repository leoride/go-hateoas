package hateoas

import ()

// This is a returned type
// Error type for REST hateoas
type Error struct {
	Status           int    `json:"status"`
	Code             int    `json:"code"`
	Property         string `json:"property,omitempty"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
	MoreInfo         string `json:"moreInfo,omitempty"`
}

func (e *Error) Error() string {
	return e.DeveloperMessage
}

// URL type for representing HREF links
type Url string

// This is a returned type
// Page type for REST hateoas
type Page struct {
	Href     Url
	Offset   int
	Limit    int
	First    Url
	Previous Url
	Next     Url
	Last     Url
	Items    []Resource
}

// PageOpts type for page options extracted from the GET parameters
type PageOpts struct {
	Offset int
	Limit  int
}

// Abstract interface for a REST resource
type Resource interface {
}

// Abstract interface for a REST resource handler
type ResourceHandler interface {
	ResourceName() string

	GetOne(string) (Resource, *Error)
	GetAll(PageOpts) ([]Resource, *Error)
	Create(Resource) (string, *Error)
	Update(string, Resource) (Resource, *Error)
	Delete(string) *Error
}
