package hateoas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handleGetAll(w http.ResponseWriter, r *http.Request, rh ResourceHandler) *Error {
	var page Page
	var err *Error

	pageOpts := PageOpts{}
	pageOpts.Limit = 20
	pageOpts.Offset = 0

	limit := r.URL.Query().Get("page.limit")
	offset := r.URL.Query().Get("page.offset")

	if limit != "" {
		limitI, err := strconv.ParseInt(limit, 10, 0)

		if err == nil {
			pageOpts.Limit = int(limitI)
		}
	}

	if offset != "" {
		offsetI, err := strconv.ParseInt(offset, 10, 0)

		if err == nil {
			pageOpts.Offset = int(offsetI)
		}
	}

	var resources []Resource
	resources, err = rh.GetAll(pageOpts)

	if err == nil {
		first := Url("http://" + r.Host + r.URL.Path + "?page.offset=0&page.limit=" + fmt.Sprint(pageOpts.Limit))

		page.Items = resources
		page.Href = Url("http://" + r.Host + r.RequestURI)
		page.Offset = pageOpts.Offset
		page.Limit = pageOpts.Limit
		page.First = &first

		json, errJ := json.MarshalIndent(page, "", "    ")

		if errJ != nil {
			err = &Error{}
			err.Status = 500
			err.Code = 1
			err.Message = "An unknown error has occurred while trying to retrieve records."
			err.DeveloperMessage = "Internal error returned by GetAll method in resource handler."
		} else {
			jsonS := string(json)
			jsonS = strings.Replace(jsonS, "\\u0026", "&", -1)

			w.WriteHeader(200)
			fmt.Fprint(w, jsonS)
		}
	}

	return err
}
