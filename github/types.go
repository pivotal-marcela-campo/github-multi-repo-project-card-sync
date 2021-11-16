package github

import (
	"encoding/json"

	"github.com/shurcooL/githubv4"
)

type PullRequest struct {
	ID      string
	URL     string
	IsDraft bool
	Author  struct {
		Login string
	}
}

type Project struct {
	ID     githubv4.ID
	Fields ProjectFields
}

type ProjectFields []ProjectField

type ProjectField struct {
	ID       githubv4.ID
	Name     string
	Settings string
}

type FieldOption struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProjectItem struct {
	ID githubv4.ID
}

func (pf ProjectFields) FindByName(name string) (ProjectField, bool) {
	for _, field := range pf {
		if field.Name == name {
			return field, true
		}
	}
	return ProjectField{}, false
}

func (pf ProjectField) FindOptionByName(name string) (FieldOption, bool) {
	var settings struct {
		Options []FieldOption `json:"options"`
	}

	json.Unmarshal([]byte(pf.Settings), &settings)

	for _, option := range settings.Options {
		if option.Name == name {
			return option, true
		}
	}
	return FieldOption{}, false
}
