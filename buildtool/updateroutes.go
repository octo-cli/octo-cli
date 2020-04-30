package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const routesTimestampFormat = "20060102T150405Z0700"

func updateRoutes(routesURL, routesPath string) error {
	resp, err := http.Get(routesURL) //nolint:gosec
	if err != nil {
		return errors.Wrapf(err, "failed getting %q", routesURL)
	}

	err = writeRoutesJSON(routesPath, resp)
	if err != nil {
		return errors.Wrap(err, "failed writing routesJSON")
	}

	err = updateLastModified(routesPath, resp)
	return errors.Wrap(err, "failed updateing last modified")
}

func writeRoutesJSON(routesPath string, resp *http.Response) error {
	outFile, err := os.Create(routesPath)
	if err != nil {
		return errors.Wrapf(err, "failed creating file %q", routesPath)
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed writing to file %q", routesPath)
	}
	err = resp.Body.Close()
	if err != nil {
		return errors.Wrap(err, "failed closing response body")
	}
	return errors.Wrapf(err, "failed closing file %q", outFile.Close())
}

func updateLastModified(routesPath string, resp *http.Response) error {
	extension := filepath.Ext(routesPath)
	if extension == "" {
		return errors.Errorf("routesPath must have a file extension (preferable .json)")
	}
	lmPath := strings.TrimSuffix(routesPath, extension) + "-last-modified.txt"
	lmHeader := resp.Header.Get("last-modified")
	lmTime, err := time.Parse(time.RFC1123, lmHeader)
	if err != nil {
		return errors.Wrapf(err, "couldn't parse last-modified time %q", lmHeader)
	}
	lmFile, err := os.Create(lmPath)
	if err != nil {
		return errors.Wrapf(err, "failed creating file %q", lmPath)
	}
	_, err = lmFile.WriteString(lmTime.Format(routesTimestampFormat))
	if err != nil {
		return errors.Wrapf(err, "failed writing to file %q", lmPath)
	}
	return errors.Wrapf(lmFile.Close(), "failed closing file %q", lmPath)
}
