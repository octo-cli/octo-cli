package model

type Endpoint struct {
	GithubCloudOnly bool
	EnabledForApps  bool
	Legacy          bool
	Deprecated      bool
	JSONBodySchema  *ParamSchema
	ID              string
	Path            string
	Method          string
	Name            string
	Concern         string
	DocsURL         string
	Summary         string
	HelpText        string
	PathParams      Params
	QueryParams     Params
	Headers         Params
	Previews        []Preview
}

type Preview struct {
	Required bool
	Name     string
	Note     string
}

type ParamType int

const (
	ParamTypeInvalid ParamType = iota
	ParamTypeString
	ParamTypeInt
	ParamTypeBool
	ParamTypeObject
	ParamTypeArray
)

type ParamSchema struct {
	Type         ParamType
	ItemSchema   *ParamSchema
	ObjectParams []Param
}

func (p *ParamSchema) Clone() *ParamSchema {
	result := ParamSchema{
		Type: p.Type,
	}
	if p.ObjectParams != nil {
		result.ObjectParams = make([]Param, len(p.ObjectParams))
		for i, objectParam := range p.ObjectParams {
			result.ObjectParams[i] = objectParam.Clone()
		}
	}
	if p.ItemSchema != nil {
		result.ItemSchema = p.ItemSchema.Clone()
	}
	return &result
}

type Param struct {
	Required bool
	Name     string
	HelpText string
	Schema   *ParamSchema
}

func (p Param) Clone() Param {
	q := p
	q.Schema = q.Schema.Clone()
	return q
}

type Params []Param

func (p Params) Clone() Params {
	if p == nil {
		return nil
	}
	result := make(Params, len(p))
	for i, param := range p {
		result[i] = param.Clone()
	}
	return result
}
