package routeparser

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/WillAbides/go-github-cli/generator/internal"
	"github.com/fatih/camelcase"
	"github.com/pkg/errors"
)

type (
	//RoutesFile represents the parsed contents of routes.json
	//routes are divided up by service such as "issues", "repos", and "git"
	RoutesFile map[string]Routes

	//Route represents one route from routes.json.  Such as "issues edit"
	Route struct {
		Description      string
		Method           string
		Path             string
		Name             string
		EnabledForApps   bool
		IDName           string
		DocumentationURL string
		Params           []*Param
		Requests         []interface{}
		Responses        []struct {
			Headers struct {
				Status      string
				ContentType string `json:"content-type"`
			}
			Body        interface{}
			Description interface{}
		}
	}

	//Param represents one parameter for a Route such as repo name or issue number
	Param struct {
		Name        string
		Type        string
		Description string
		Default     interface{}
		Required    bool
		Enum        []interface{}
		Location    string
		MapTo       string
	}

	//Routes is a collection or Routes
	Routes []*Route
)

//ParseRoutesFile parses the routes.json at a given path and returns a *RoutesFile
func ParseRoutesFile(file string) (map[string]Routes, error) {
	var sm map[string]Routes
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed reading file %q", file)
	}
	err = json.Unmarshal(bts, &sm)
	return sm, errors.Wrap(err, "failed unmarshalling json")
}

//FindByIDName returns the Route with the given IDName
func (r *Routes) FindByIDName(idName string) *Route {
	for _, route := range *r {
		if route != nil && route.IDName == idName {
			return route
		}
	}
	return nil
}

//FieldParam returns the parameter for a given field
func (r *Route) FieldParam(field reflect.StructField) *Param {
	jsonTag := field.Tag.Get("json")
	var name string

	if jsonTag == "" {
		name = strings.ToLower(strings.Join(camelcase.Split(field.Name), "-"))
	} else {
		name = strings.Split(jsonTag, ",")[0]
	}

	for _, param := range r.Params {
		if param != nil && strings.ToLower(param.Name) == name {
			return param
		}
	}
	return nil
}

//ArgParam returns the Param for the given argument name
func (r *Route) ArgParam(name string) *Param {
	for _, param := range r.Params {
		if param != nil && internal.ToArgName(param.Name) == internal.ToArgName(name) {
			return param
		}
	}
	return nil
}
