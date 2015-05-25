package main

import (
	"github.com/site-builder/worker/cloner"
	"github.com/site-builder/worker/directory_helper"
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/runner"
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
