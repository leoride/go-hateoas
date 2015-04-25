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

	http.HandleFunc(location, errorWrapper(rootPathLength, rh))

	return err
}

func errorWrapper(rootPathLength int, rh ResourceHandler) func(w http.ResponseWriter, r *http.Request) {
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
			err = handle(w, r, rh, id)
		}

		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(err.Status)

			errJson, _ := json.MarshalIndent(err, "", "    ")
			fmt.Fprint(w, string(errJson))
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request, rh ResourceHandler, id string) *Error {
	var err *Error
	fmt.Println("Request received", r.RequestURI)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case "GET":
		if id == "" {
			fmt.Println("GET request (all)")
			err = handleGetAll(w, r, rh)
		} else {
			fmt.Println("GET request (single resource)", id)
			err = handleGetOne(w, r, rh, id)
		}
		break

	case "POST":
		if id == "" {
			fmt.Println("POST request (create)")
			err = handleCreate(w, r, rh)
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
			err = handleUpdate(w, r, rh, id)
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
			err = handleDelete(w, r, rh, id)
		}
		break
	}

	return err
}
