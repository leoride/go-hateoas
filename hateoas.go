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
			w.WriteHeader(err.Status)

			errJson, _ := json.MarshalIndent(err, "", "    ")
			fmt.Fprint(w, string(errJson))
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request, rh ResourceHandler, id string) *Error {
	var err *Error

    //TODO: remove this block
	err = &Error{}
	err.Status = 500
	err.Code = 1
	err.Message = "Under construction. Please check again later."
	err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

    fmt.Println("Request received", r.RequestURI)

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

func handleCreate(w http.ResponseWriter, r *http.Request, rh ResourceHandler) *Error {
    var err *Error

    err = &Error{}
    err.Status = 500
    err.Code = 1
    err.Message = "Create is under construction. Please check again later."
    err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

    return err
}

func handleUpdate(w http.ResponseWriter, r *http.Request, rh ResourceHandler, id string) *Error {
    var err *Error

    err = &Error{}
    err.Status = 500
    err.Code = 1
    err.Message = "Update is under construction. Please check again later."
    err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

    return err
}

func handleDelete(w http.ResponseWriter, r *http.Request, rh ResourceHandler, id string) *Error {
    var err *Error

    err = &Error{}
    err.Status = 500
    err.Code = 1
    err.Message = "Delete is under construction. Please check again later."
    err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

    return err
}

func handleGetAll(w http.ResponseWriter, r *http.Request, rh ResourceHandler) *Error {
    var err *Error

    err = &Error{}
    err.Status = 500
    err.Code = 1
    err.Message = "Get all is under construction. Please check again later."
    err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

    return err
}

func handleGetOne(w http.ResponseWriter, r *http.Request, rh ResourceHandler, id string) *Error {
    var err *Error

    err = &Error{}
    err.Status = 500
    err.Code = 1
    err.Message = "Get one is under construction. Please check again later."
    err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

    return err
}