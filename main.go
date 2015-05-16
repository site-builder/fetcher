package main

import (
	"github.com/site-builder/fetcher/cloner"
	"github.com/site-builder/fetcher/locator"
	"github.com/site-builder/fetcher/logger"
	"github.com/site-builder/fetcher/randomizer"
	"github.com/site-builder/fetcher/runner"
)

var log = logger.CreateLogger("main")

func createTempDirectory() locator.Locator {
	randomizer := randomizer.CreateRandomizer()
	directory := locator.NewTempDirectoryLocator(randomizer)
	runner := runner.NewRunner()

	log.Info("Creating temporary directory %s", directory.Location())

	if err := runner.Run("mkdir", "-p", directory.Location()); err != nil {
		log.Error("Error creating directory %s", err)
	}

	return directory
}

func removeTempDirectory(directory locator.Locator) {
	runner := runner.NewRunner()

	log.Info("Removing temporary directory %s", directory.Location())

	if err := runner.Run("rm", "-rf", directory.Location()); err != nil {
		log.Error("Error removing directory %s", err)
	}
}

type tempDirectoryCallback func(l locator.Locator)

func withTemporaryDirectory(fn tempDirectoryCallback) {
	temporaryDirectory := createTempDirectory()

	fn(temporaryDirectory)

	removeTempDirectory(temporaryDirectory)
}

func main() {
	runner := runner.NewRunner()
	source := locator.NewGitLocator("/tmp/frontend")
	// destination := locator.NewGitLocator("git@github.com:site-builder/frontend.git")

	withTemporaryDirectory(func(temp locator.Locator) {
		cloner.Clone(runner, source, temp, logger.CreateLogger("cloner"))
		// builder.Build(runner, temp, logger.CreateLogger("builder"))
		// deployer.Deploy(runner, temp, logger.CreateLogger("deployer"))
	})
}
