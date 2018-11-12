package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func getRoutesContent(tagName string) (string, error) {
	fileSha, err := runOcto("git", "get-tree", "--owner", "octokit", "--repo", "routes",
		"--tree_sha", tagName,
		"--format", `{{range .tree}}{{if eq .path "index.json"}}{{.sha}}{{end}}{{end}}`)
	if err != nil {
		return "", err
	}
	encoded, err := runOcto("git", "get-blob", "--owner", "octokit", "--repo", "routes",
		"--file_sha", fileSha, "--format", "{{.content}}")
	if err != nil {
		return "", err
	}
	routesBytes, err := base64.StdEncoding.DecodeString(encoded)
	routesContent := string(routesBytes)
	return routesContent, err
}

func updateRoutes(routesPath, githubToken string) error {
	err := os.Setenv("GITHUB_TOKEN", githubToken)
	if err != nil {
		return errors.Wrap(err, "failed setting env GITHUB_TOKEN")
	}
	tag, err := runOcto("repos", "get-latest-release", "--owner", "octokit", "--repo", "routes",
		"--format", "{{.tag_name}}")
	if err != nil {
		return errors.Wrap(err, "failed getting latest release for octokit/routes")
	}
	routesJSON, err := getRoutesContent(tag)
	if err != nil {
		return errors.Wrap(err, "failed getting contents of index.json")
	}
	err = ioutil.WriteFile(routesPath, []byte(routesJSON), 0644)
	if err != nil {
		return errors.Wrapf(err, "failed writing to %q", routesPath)
	}
	routesTagPath := routesPath + "-tag.txt"
	err = ioutil.WriteFile(routesTagPath, []byte(tag), 0644)
	return errors.Wrapf(err, "failed writing to %q", routesTagPath)
}
