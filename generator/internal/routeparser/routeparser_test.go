package routeparser_test

import (
	"reflect"
	"testing"

	. "github.com/WillAbides/go-github-cli/generator/internal/routeparser"
	"github.com/stretchr/testify/assert"
)

func TestParseRoutesFile(t *testing.T) {
	routesMap, err := ParseRoutesFile("../../testdata/routes.json")
	assert.Nil(t, err)
	var mapKeys []string
	for key := range routesMap {
		mapKeys = append(mapKeys, key)
	}
	assert.Contains(t, mapKeys, "issues")
}

func TestRoutes_FindByIdName(t *testing.T) {
	want := &Route{IDName: "foo"}
	routes := Routes{
		nil,
		&Route{IDName: "other"},
		want,
	}
	got := routes.FindByIdName("foo")
	assert.Equal(t, want, got)
}

func TestRoute(t *testing.T) {
	route := &Route{
		Params: []*Param{
			nil,
			{},
			{Name: "foo-bar", Description: "BazQux"},
		},
	}

	t.Run("FieldParam", func(t *testing.T) {

		t.Run("without json tag", func(t *testing.T) {
			sf := reflect.StructField{
				Name: "FooBar",
			}
			got := route.FieldParam(sf)
			assert.Equal(t, "BazQux", got.Description)
		})

		t.Run("with a json tag", func(t *testing.T) {
			sf := reflect.StructField{
				Name: "Whatever",
				Tag:  reflect.StructTag(`json:"foo-bar"`),
			}
			got := route.FieldParam(sf)
			assert.Equal(t, "BazQux", got.Description)
		})
	})

	t.Run("ArgParam", func(t *testing.T) {
		got := route.ArgParam("FooBar")
		assert.Equal(t, "BazQux", got.Description)
	})
}

func TestRoute_FieldParam(t *testing.T) {
	route := &Route{
		Params: []*Param{
			nil,
			{},
			{Name: "foo-bar", Description: "BazQux"},
		},
	}

	t.Run("without json tag", func(t *testing.T) {
		sf := reflect.StructField{
			Name: "FooBar",
		}
		got := route.FieldParam(sf)
		assert.Equal(t, "BazQux", got.Description)
	})

	t.Run("with a json tag", func(t *testing.T) {
		sf := reflect.StructField{
			Name: "Whatever",
			Tag:  reflect.StructTag(`json:"foo-bar"`),
		}
		got := route.FieldParam(sf)
		assert.Equal(t, "BazQux", got.Description)
	})
}
