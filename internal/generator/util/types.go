package util

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

var ParamTypes = map[string]string{
	"string":    "string",
	"integer":   "int64",
	"string[]":  "[]string",
	"integer[]": "[]int64",
	"boolean":   "bool",
	"object":    "internal.JSONObject",
	"object[]":  "[]internal.JSONObject",
}

func GetPropType(schema *openapi3.Schema) string {
	if schema == nil {
		return ""
	}
	if schema.Type != "array" || schema.Items == nil {
		return schema.Type
	}
	itemType := schema.Items.Value.Type
	return fmt.Sprintf("%s[]", itemType)
}

//ToArgName takes input like "foo-bar" and returns "FooBar"
func ToArgName(in string) string {
	out := in
	for _, separator := range []string{"_", "-", "."} {
		words := strings.Split(out, separator)
		for i, word := range words {
			words[i] = strings.Title(word)
		}
		out = strings.Join(words, "")
	}
	return out
}
