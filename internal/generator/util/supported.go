package util

import (
	"github.com/octo-cli/octo-cli/internal/model"
)

func EndpointIsUnsupported(endpoint model.Endpoint) bool {
	bodySchema := endpoint.JSONBodySchema
	if bodySchema == nil {
		return false
	}
	for _, param := range bodySchema.ObjectParams {
		if !param.Required {
			continue
		}
		if !IsSupportedModelParam(param) {
			return true
		}
	}
	return false
}

func IsSupportedModelParam(param model.Param) bool {
	schema := param.Schema
	if schema == nil {
		return false
	}
	if schema.Type == model.ParamTypeInvalid {
		return false
	}
	if schema.ItemSchema != nil && schema.ItemSchema.Type == model.ParamTypeInvalid {
		return false
	}
	for _, objectParam := range schema.ObjectParams {
		ok := IsSupportedModelParam(objectParam)
		if !ok {
			return false
		}
	}
	return true
}
