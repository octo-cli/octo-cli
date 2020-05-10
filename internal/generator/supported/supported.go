package supported

import (
	"sort"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/octo-cli/octo-cli/internal/generator/swaggerparser"
	"github.com/octo-cli/octo-cli/internal/generator/util"
)

func OperationIsUnsupported(op *openapi3.Operation, routePath, method string) bool {
	for _, check := range unsupportedChecks {
		if check.check(op, routePath, method) {
			return true
		}
	}
	return false
}

type unsupportedCheck struct {
	check  func(op *openapi3.Operation, routePath, method string) bool
	reason string
}

var unsupportedChecks = []unsupportedCheck{
	{
		reason: "manually set to unsupported",
		check: func(op *openapi3.Operation, routePath, method string) bool {
			disList := []string{
				"markdown/render-raw",
			}
			for _, dissed := range disList {
				if op.OperationID == dissed {
					return true
				}
			}
			return false
		},
	},
	{
		reason: "contains required body parameters of an unsupported type",
		check: func(op *openapi3.Operation, _, _ string) bool {
			if op.RequestBody == nil || op.RequestBody.Value.Content.Get("application/json") == nil {
				return false
			}
			bodySchema := op.RequestBody.Value.Content.Get("application/json").Schema.Value
			required := map[string]bool{}
			for _, s := range bodySchema.Required {
				required[s] = true
			}
			for name, schemaRef := range bodySchema.Properties {
				if !required[name] {
					continue
				}
				if !IsSupportedParam(schemaRef) {
					return true
				}
			}
			return false
		},
	},
	{
		reason: "accepts non-json body",
		check: func(op *openapi3.Operation, routePath, method string) bool {
			if op.RequestBody == nil || op.RequestBody.Value == nil {
				return false
			}
			for tp := range op.RequestBody.Value.Content {
				if tp != "application/json" {
					return true
				}
			}
			return false
		},
	},
	{
		reason: "invalid response code",
		check: func(op *openapi3.Operation, routePath, method string) bool {
			for code := range op.Responses {
				_, err := strconv.Atoi(code)
				if err != nil {
					return true
				}
			}
			return false
		},
	},
	{
		reason: "unhandled server in openapi spec",
		check: func(op *openapi3.Operation, routePath, method string) bool {
			return op.Servers != nil
		},
	},
}

func isSupportedParamType(tp string) bool {
	_, ok := util.ParamTypes[tp]
	return ok
}

func IsSupportedParam(ref *openapi3.SchemaRef) bool {
	if ref == nil || ref.Value == nil {
		return false
	}
	tp := util.GetPropType(ref.Value)
	if !isSupportedParamType(tp) {
		return false
	}
	if ref.Value.Type != "object" {
		return true
	}
	for _, subRef := range ref.Value.Properties {
		if subRef == nil || subRef.Value == nil {
			return false
		}
		subTp := util.GetPropType(subRef.Value)
		if !isSupportedParamType(subTp) {
			return false
		}
	}
	return true
}

type OpInfo struct {
	Method string
	Path   string
	Op     *openapi3.Operation
}

func SvcOperations(swagger *openapi3.Swagger, svcName string) []OpInfo {
	var result []OpInfo
	swaggerparser.ForEachOperation(swagger, func(path, method string, op *openapi3.Operation) {
		if swaggerparser.GetOperationSvcName(op) != svcName {
			return
		}
		if OperationIsUnsupported(op, path, method) {
			return
		}
		result = append(result, OpInfo{
			Method: method,
			Path:   path,
			Op:     op,
		})
	})
	return result
}

func RefFilter(ref *openapi3.SchemaRef) bool {
	return IsSupportedParam(ref)
}

type UnsupportedOperation struct {
	Operation *openapi3.Operation
	RoutePath string
	Method    string
	Reason    string
}

func GetUnsupportedOperations(swagger *openapi3.Swagger) []UnsupportedOperation {
	var result []UnsupportedOperation
	//nolint:errcheck // err is always nil
	swaggerparser.ForEachOperation(swagger, func(path, method string, op *openapi3.Operation) {
		for _, check := range unsupportedChecks {
			if check.check(op, path, method) {
				result = append(result, UnsupportedOperation{
					Operation: op,
					RoutePath: path,
					Method:    method,
					Reason:    check.reason,
				})
				break
			}
		}
	})
	sort.Slice(result, func(i, j int) bool {
		return result[i].Operation.OperationID < result[j].Operation.OperationID
	})
	return result
}
