package internal

import (
	"github.com/fatih/camelcase"
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
