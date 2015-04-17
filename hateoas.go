package hateoas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle(apiPath string, rh ResourceHandler) error {
	var err error

	location := fmt.Sprint(apiPath, "/", rh.ResourceName())
	http.HandleFunc(location, errorWrapper)

	return err
}

func errorWrapper(w http.ResponseWriter, r *http.Request) {
	var err *Error

	err = handle(w, r)

	if err != nil {
		w.WriteHeader(err.Status)

		errJson, _ := json.MarshalIndent(err, "", "    ")
		fmt.Fprint(w, string(errJson))
	}
}

func handle(w http.ResponseWriter, r *http.Request) *Error {
	var err *Error

	err = &Error{}
	err.Status = 500
	err.Code = 1
	err.Message = "Under construction. Please check again later."
	err.DeveloperMessage = "API is not ready yet. Please contact tomg@leoride.com for more information."

	return err
}
