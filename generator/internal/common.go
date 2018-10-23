package internal

import (
	"github.com/fatih/camelcase"
	"github.com/fatih/structtag"
	"strings"
)

//ToArgName takes input like "foo-bar" and returns "FooBar"
func ToArgName(in string) string {
	out := in
	for _, separator := range []string{"_", "-"} {
		words := strings.Split(out, separator)
		for i, word := range words {
			words[i] = strings.Title(word)
		}
		out = strings.Join(words, "")
	}
	return out
}

//Unexport takes some camelcased strings and returns a camelcased, unexported, truncated version
func Unexport(name ...string) string {
	var words []string
	for _, v := range name {
		w := camelcase.Split(v)
		for i, ww := range w {
			w[i] = strings.Title(ww)
		}
		words = append(words, w...)
	}

	if len(words) < 1 {
		return ""
	}
	words[0] = strings.ToLower(words[0])
	return strings.Join(words, "")
}

//RunMethodArg is an argument for a RunMethod
type RunMethodArg struct {
	Name  string
	IsPtr bool
}

//RunMethod represents the Run() method for a cmd struct
type RunMethod struct {
	StructName string
	HasElement bool
	SvcName    string
	FuncName   string
	Args       []RunMethodArg
}

// StructTmplHelper represents a struct for the template to build. It can also include other structs and a run method
//    when it represents a command struct
type StructTmplHelper struct {
	Name         string
	Fields       []StructField
	ChildStructs []StructTmplHelper
	RunMethod    *RunMethod
	ToFunc       *ToFunc
}

// StructField is one field in a StructTmplHelper
type StructField struct {
	Name string
	Type string
	Tags *structtag.Tags
}

//ToFunc represents the function that converts a cli options struct to a go-github options struct
//  an example is from issues create is:
//    func (t issuesCreateCmdIssueRequestFlags) toIssueRequest(k *kong.Context) *github.IssueRequest
type ToFunc struct {
	ReceiverName            string
	TargetName              string
	TargetType              string
	ValSetters              []ValSetter
}

//IncludePointerHelper determines whether the generated func should include the "isValueSet" helper
func (t *ToFunc) IncludePointerHelper() bool {
	for _, v := range t.ValSetters {
		if v.TargetIsPtr {
			return true
		}
	}
	return false
}

//ValSetter sets one value in a toFunc
//  example output: `val.LockReason = t.LockReason`
//
//  or:
//    if isValueSet("labels") {
//      val.Labels = &t.Labels
//    }
type ValSetter struct {
	TargetIsPtr bool
	Name        string
	FlagName    string
}

// ToPkg represents the go package that will be created for a svc
type Pkg struct {
	PackageName string
	Imports     []string
	Structs     []*StructTmplHelper
}
