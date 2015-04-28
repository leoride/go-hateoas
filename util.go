package hateoas

import (
	"fmt"
	"reflect"
)

type hateoasResource map[string]interface{}

func toHateoasResources(r []Resource, baseUrl string) []hateoasResource {
	hateoasResources := make([]hateoasResource, len(r))

	for index, value := range r {
		hateoasResources[index] = toHateoasResource(value, baseUrl)
	}

	return hateoasResources
}

func toHateoasResource(r Resource, baseUrl string) hateoasResource {
	resourceId := r.GetId()
	resourceUrl := Url(fmt.Sprint(baseUrl + resourceId))

	hateoasResource := hateoasResource{}
	hateoasResource["Href"] = resourceUrl

	v := reflect.ValueOf(r)
	vi := reflect.Indirect(v)
	for i := 0; i < v.NumField(); i++ {
		fName := vi.Type().Field(i).Name
		fValue := v.Field(i).Interface()
		hateoasResource[fName] = fValue
	}

	return hateoasResource
}
