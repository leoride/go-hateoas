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
	var count int

	resources, err = rh.GetAll(pageOpts)

	if err == nil {
		count, err = rh.Count()
	}

	if err == nil {
		first := getFirstUrl(r, pageOpts.Limit)
		last := getLastUrl(r, pageOpts.Limit, pageOpts.Offset, count)
		prev := getPrevUrl(r, pageOpts.Limit, pageOpts.Offset, count)
		next := getNextUrl(r, pageOpts.Limit, pageOpts.Offset, count)

		baseUrl := "http://" + r.Host + r.URL.Path
		baseUrl = strings.Replace(baseUrl, rh.ResourceName()+"/", "", -1)

		page.Items = toHateoasResources(resources, baseUrl, rh.ResourceName())
		page.Href = getSelfUrl(r)
		page.Offset = pageOpts.Offset
		page.Limit = pageOpts.Limit
		page.First = first
		page.Last = last
		page.Previous = prev
		page.Next = next
		page.TotalItems = count

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

func getSelfUrl(r *http.Request) Url {
	var url Url
	url = Url("http://" + r.Host + r.RequestURI)
	return url
}

func getFirstUrl(r *http.Request, limit int) *Url {
	var url Url
	url = Url("http://" + r.Host + r.URL.Path + "?page.offset=0&page.limit=" + fmt.Sprint(limit))
	return &url
}

func getLastUrl(r *http.Request, limit int, offset int, count int) *Url {
	var url Url
	newOffset := (count / limit) * limit

	if newOffset == count {
		newOffset -= limit
	}

	url = Url("http://" + r.Host + r.URL.Path + "?page.offset=" + fmt.Sprint(newOffset) + "&page.limit=" + fmt.Sprint(limit))

	return &url
}

func getPrevUrl(r *http.Request, limit int, offset int, count int) *Url {
	var url Url

	if (offset - limit) >= 0 {
		newOffset := offset - limit
		url = Url("http://" + r.Host + r.URL.Path + "?page.offset=" + fmt.Sprint(newOffset) + "&page.limit=" + fmt.Sprint(limit))
		return &url
	} else {
		return nil
	}
}

func getNextUrl(r *http.Request, limit int, offset int, count int) *Url {
	var url Url

	if (offset + limit) < count {
		newOffset := offset + limit
		url = Url("http://" + r.Host + r.URL.Path + "?page.offset=" + fmt.Sprint(newOffset) + "&page.limit=" + fmt.Sprint(limit))
		return &url
	} else {
		return nil
	}
}
