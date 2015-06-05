package main

import (
	"github.com/site-builder/worker/builder"
	"github.com/site-builder/worker/cloner"
	"github.com/site-builder/worker/deployer"
	"github.com/site-builder/worker/directory_helper"
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/logger"
	"github.com/site-builder/worker/runner"

	"flag"
)

var log = logger.CreateLogger("worker")

func main() {
	var sourceRepo string
	var sourceBranch string

	var destinationRepo string
	var destinationBranch string

	flag.StringVar(&sourceRepo, "source-repo", "", "repo to clone")
	flag.StringVar(&sourceBranch, "source-branch", "master", "branch to clone")

	flag.StringVar(&destinationRepo, "destination-repo", "", "repo to send deploy")
	flag.StringVar(&destinationBranch, "destination-branch", "gh-pages", "branch to send deploy")

	flag.Parse()

	log.Info("Starting build")

	source := locator.NewGitLocatorWithBranch(sourceRepo, sourceBranch)
	destination := locator.NewGitLocatorWithBranch(destinationRepo, destinationBranch)

	runner := runner.NewRunner()
	directoryHelper := directory_helper.NewDirectoryHelper(runner)
	cloner := cloner.NewCloner(runner)
	builder := builder.NewBuilder(runner)
	deployer := deployer.NewDeployer(runner)

	directoryHelper.WithTemporaryDirectory(func(tmpSource locator.Locator) {
		directoryHelper.WithTemporaryDirectory(func(tmpDestination locator.Locator) {
			cloner.Clone(source, tmpSource)
			builder.Build(tmpSource, tmpDestination)
			deployer.Deploy(tmpDestination, destination)
		})
	})

	log.Info("Build complete")
}
