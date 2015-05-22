package main

import (
	"github.com/site-builder/fetcher/cloner"
	"github.com/site-builder/fetcher/directory_helper"
	"github.com/site-builder/fetcher/locator"
	"github.com/site-builder/fetcher/runner"
)

func main() {
	runner := runner.NewRunner()
	source := locator.NewGitLocator("/tmp/frontend")
	directoryHelper := directory_helper.NewDirectoryHelper(runner)
	// destination := locator.NewGitLocator("git@github.com:site-builder/frontend.git")

	directoryHelper.WithTemporaryDirectory(func(temp locator.Locator) {
		cloner.Clone(runner, source, temp)
		// builder.Build(runner, temp)
		// deployer.Deploy(runner, temp)
	})
}
