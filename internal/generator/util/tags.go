package util

import (
	"github.com/fatih/structtag"
)

//FieldTags returns the tags for a command field
func FieldTags(name string, required bool) *structtag.Tags {
	if required {
		return NewTags(NewTag("required", ""), NewTag("name", name))
	}
	return NewTags(NewTag("name", name))
}

//NewTag is a helper to create a new *structtag.Tag with fewer lines of code
func NewTag(key, name string, options ...string) *structtag.Tag {
	return &structtag.Tag{
		Key:     key,
		Name:    name,
		Options: options,
	}
}

//NewTags creates a new *structtag.Tags from a list of tags
//  it will panic if one of the tags has no key
func NewTags(tag ...*structtag.Tag) *structtag.Tags {
	tags := &structtag.Tags{}
	for _, tag := range tag {
		err := tags.Set(tag)
		if err != nil {
			panic(err)
		}
	}
	return tags
}

func TagsHasKey(tags *structtag.Tags, key string) bool {
	if tags == nil {
		return false
	}
	for _, s := range tags.Keys() {
		if s == key {
			return true
		}
	}
	return false
}
