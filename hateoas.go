package hateoas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Handle(apiPath string, rh ResourceHandler) error {
	var err error

	location := fmt.Sprint(apiPath, "/", rh.ResourceName(), "/")
	rootPathLength := len(location)

	http.HandleFunc(location, errorWrapper(rootPathLength))

	return err
}

func errorWrapper(rootPathLength int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err *Error

		id := r.URL.Path[rootPathLength:]

		if strings.Contains(id, "/") {
			err = &Error{}
			err.Status = 500
			err.Code = 1
			err.Message = "Invalid request. Resource could not be fetched."
			err.DeveloperMessage = "Request for resource is incorrect. More than one path parameters (ID) is not supported for a resource."
		}

		if err == nil {
			err = handle(w, r, id)
		}

		if err != nil {
			w.WriteHeader(err.Status)

			errJson, _ := json.MarshalIndent(err, "", "    ")
			fmt.Fprint(w, string(errJson))
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request, id string) *Error {
	var err *Error

	err = &Error{}
	err.Status = 500
	err.Code = 1
	err.Message = "Under construction. Please check again later."
	err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

	switch r.Method {
	case "GET":
		if id == "" {
			fmt.Println("GET request (all)")
		} else {
			fmt.Println("GET request (single resource)", id)
		}
		break

	case "POST":
		if id == "" {
			fmt.Println("POST request (create)")
		} else {
			err = &Error{}
			err.Status = 500
			err.Code = 1
			err.Message = "Invalid request. Resource could not be created."
			err.DeveloperMessage = "Request for resource is incorrect. No path parameters are needed for creating a resource."
		}
		break

	case "PUT":
		if id == "" {
			err = &Error{}
			err.Status = 500
			err.Code = 1
			err.Message = "Invalid request. Resource could not be updated."
			err.DeveloperMessage = "Request for resource is incorrect. A path parameter (ID) is needed for updating a resource."
		} else {
			fmt.Println("PUT request (update resource)", id)
		}
		break

	case "DELETE":
		if id == "" {
			err = &Error{}
			err.Status = 500
			err.Code = 1
			err.Message = "Invalid request. Resource could not be deleted."
			err.DeveloperMessage = "Request for resource is incorrect. A path parameter (ID) is needed for deleting a resource."
		} else {
			fmt.Println("DELETE request (delete resource)", id)
		}
		break
	}

	return err
}
