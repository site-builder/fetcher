package main

import (
	"github.com/site-builder/worker/builder"
	"github.com/site-builder/worker/cloner"
	"github.com/site-builder/worker/deployer"
	"github.com/site-builder/worker/directory_helper"
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/logger"
	"github.com/site-builder/worker/runner"

	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	log = logger.CreateLogger("worker")

	app = kingpin.New("worker", "Builds a Jekyll site from a git repository and pushes the results to another git repository.")

	sourceRepo   = app.Flag("source-repo", "Repo to clone.").Required().String()
	sourceBranch = app.Flag("source-branch", "Branch to clone.").Default("master").String()

	destinationRepo   = app.Flag("destination-repo", "Repo to send deploy.").Required().String()
	destinationBranch = app.Flag("destination-branch", "Branch to send deploy.").Default("gh-pages").String()
)

func main() {
	app.Parse(os.Args[1:])

	log.Info("Starting build")

	source := locator.NewGitLocatorWithBranch(*sourceRepo, *sourceBranch)
	destination := locator.NewGitLocatorWithBranch(*destinationRepo, *destinationBranch)

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
