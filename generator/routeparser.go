package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

type (
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
		Previews         []*Preview
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

	//Preview is a preview header
	Preview struct {
		Name        string
		Description string
		Required    bool
	}
)

//ParseRoutesFile parses the routes.json at a given path and returns a *RoutesFile
func ParseRoutesFile(file string) (map[string]Routes, error) {
	var sm map[string]Routes
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed reading file %q", file)
	}
	err = json.Unmarshal(bts, &sm)
	for _, rts := range sm {
		for _, rt := range rts {

			for _, param := range rt.Params {
				// This is a stupid hack to get rid of unescaped double quotes in parameter help.
				param.Description = strings.Replace(param.Description, `"`, `'`, -1)
			}
		}
	}
	return sm, errors.Wrap(err, "failed unmarshalling json")
}
