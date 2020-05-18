package util

import (
	"strings"

	"github.com/octo-cli/octo-cli/internal/model"
)

func FixPreviewNote(note string) string {
	note = strings.TrimSpace(note)
	note = strings.Split(note, "```")[0]
	note = strings.TrimSpace(note)
	setThisFlagPhrases := []string{
		"provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header",
		"provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header",
		"provide the following custom [media type](https://developer.github.com/v3/media) in the `Accept` header",
	}
	for _, phrase := range setThisFlagPhrases {
		note = strings.ReplaceAll(note, phrase, "set this flag")
	}
	note = strings.TrimSpace(note)
	note = strings.TrimSuffix(note, ":")
	note = strings.TrimSpace(note)
	note = strings.TrimSuffix(note, ".") + "."
	return note
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

var schemaParamTypes = map[model.ParamType]string{
	model.ParamTypeString: "string",
	model.ParamTypeInt:    "int64",
	model.ParamTypeBool:   "bool",
	model.ParamTypeObject: "internal.JSONObject",
}

func SchemaParamType(schema *model.ParamSchema) string {
	if schema == nil {
		return ""
	}
	if schema.Type == model.ParamTypeArray {
		return "[]" + SchemaParamType(schema.ItemSchema)
	}
	return schemaParamTypes[schema.Type]
}
