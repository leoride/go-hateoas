package hateoas

import (
	"fmt"
	"reflect"
	"strings"
)

type hateoasResource map[string]interface{}

func toHateoasResources(r []Resource, baseUrl string, resourceName string) []hateoasResource {
	hateoasResources := make([]hateoasResource, len(r))

	for index, value := range r {
		hateoasResources[index] = toHateoasResource(value, baseUrl, resourceName)
	}

	return hateoasResources
}

func toHateoasResource(r Resource, baseUrl string, resourceName string) hateoasResource {
	resourceId := r.GetId()
	resourceUrl := Url(fmt.Sprint(baseUrl + resourceName + "/" + resourceId))

	hateoasResource := hateoasResource{}
	hateoasResource["Href"] = resourceUrl

	v := reflect.ValueOf(r)
	vi := reflect.Indirect(v)

	for i := 0; i < v.NumField(); i++ {
		fName := vi.Type().Field(i).Name
		fValue := v.Field(i).Interface()

		resource, isResource := fValue.(Resource)

		if isResource {
			if resource.GetId() != "" {
				hateoasResource[fName] = toHateoasResource(resource, baseUrl, strings.ToLower(fName))
			} else {
				hateoasResource[fName] = nil
			}
		} else {
			hateoasResource[fName] = fValue
		}
	}

	return hateoasResource
}
