package main

import (
	"github.com/site-builder/fetcher/cloner"
	"github.com/site-builder/fetcher/commander"
	"github.com/site-builder/fetcher/locator"
	"github.com/site-builder/fetcher/logger"
	"github.com/site-builder/fetcher/randomizer"
)

var log = logger.CreateLogger("main")

func createTempDirectory() locator.Locator {
	randomizer := randomizer.CreateRandomizer()
	locator := locator.NewTempDirectoryLocator(randomizer)

	command := commander.CreateCommand("mkdir", "-p", locator.Location())

	if err := command.Run(); err != nil {
		log.Error("Error creating directory %s", err)
	}

	return locator
}

func removeTempDirectory(directory locator.Locator) {
	command := commander.CreateCommand("rm", "-rf", directory.Location())

	if err := command.Run(); err != nil {
		log.Error("Error removing directory %s", err)
	}
}

type tempDirectoryCallback func(l locator.Locator)

func withTemporaryDirectory(fn tempDirectoryCallback) {
	temporaryDirectory := createTempDirectory()

	log.Info("Created temporary directory %s", temporaryDirectory.Location())

	fn(temporaryDirectory)

	removeTempDirectory(temporaryDirectory)

	log.Info("Removed temporary directory %s", temporaryDirectory.Location())
}

func main() {
	source := locator.NewGitLocator("/tmp/frontend")
	// destination := locator.NewGitLocator("git@github.com:site-builder/frontend.git")

	withTemporaryDirectory(func(temp locator.Locator) {
		log.Info("About to clone from: %s", source.Location())

		cloner.Clone(source, temp, logger.CreateLogger("cloner"))
		// builder.Build(temp, logger.CreateLogger("builder"))
		// deployer.Deploy(temp, destination, logger.CreateLogger("deployer"))
	})
}
