package directory_helper

import (
	"github.com/site-builder/fetcher/locator"
	"github.com/site-builder/fetcher/runner"
)

type callback func(locator.Locator)

type DirectoryHelper interface {
	WithTemporaryDirectory(callback)
}

type directoryHelper struct {
	runner runner.Runner
}

func NewDirectoryHelper() DirectoryHelper {
	return &directoryHelper{}
}

func (directoryHelper *directoryHelper) WithTemporaryDirectory(callback callback) {
}
