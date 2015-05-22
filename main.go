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

	directoryHelper.WithTemporaryDirectory(func(tmp locator.Locator) {
		cloner.Clone(runner, source, tmp)
		// builder.Build(runner, tmp)
		// deployer.Deploy(runner, tmp)
	})
}
