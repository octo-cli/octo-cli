package util

import (
	"sort"
	"strings"

	"github.com/octo-cli/octo-cli/internal/model"
)

func RemoveOwnerParams(endpoints []model.Endpoint) {
	for i, endpoint := range endpoints {
		ownerIdx := -1
		for i, param := range endpoint.PathParams {
			if param.Name == "owner" {
				ownerIdx = i
				break
			}
		}
		if ownerIdx == -1 || ownerIdx == len(endpoint.PathParams)-1 {
			continue
		}
		repoIdx := ownerIdx + 1
		if endpoint.PathParams[repoIdx].Name != "repo" {
			continue
		}
		endpoint.PathParams[repoIdx].HelpText = "repository in OWNER/REPO form"
		endpoint.PathParams = append(endpoint.PathParams[:ownerIdx], endpoint.PathParams[ownerIdx+1:]...)
		endpoint.Path = strings.Replace(endpoint.Path, `/{owner}/{repo}`, `/{repo}`, 1)
		endpoints[i] = endpoint
	}
}

func AllConcerns(endpoints []model.Endpoint) []string {
	concerns := make(map[string]bool, len(endpoints))
	for _, endpoint := range endpoints {
		concerns[endpoint.Concern] = true
	}
	result := make([]string, 0, len(concerns))
	for c := range concerns {
		result = append(result, c)
	}
	sort.Strings(result)
	return result
}

func flattenParams(params model.Params, namePrefix string, parentRequired bool) model.Params {
	result := make(model.Params, 0, len(params))
	for _, param := range params {
		param = param.Clone()
		if !parentRequired {
			param.Required = false
		}
		param.Name = namePrefix + param.Name
		if param.Schema.Type != model.ParamTypeObject {
			result = append(result, param)
			continue
		}
		moreParams := flattenParams(param.Schema.ObjectParams, param.Name+".", param.Required)
		if len(moreParams) == 0 {
			result = append(result, param)
			continue
		}
		result = append(result, moreParams...)
	}
	return result
}

func FlattenParams(params model.Params) model.Params {
	return flattenParams(params, "", true)
}
