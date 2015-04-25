package hateoas

import "net/http"

func handleDelete(w http.ResponseWriter, r *http.Request, rh ResourceHandler, id string) *Error {
	var err *Error

	err = &Error{}
	err.Status = 500
	err.Code = 1
	err.Message = "Delete is under construction. Please check again later."
	err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

	return err
}
