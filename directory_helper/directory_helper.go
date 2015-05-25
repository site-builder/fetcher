package directory_helper

import (
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/logger"
	"github.com/site-builder/worker/randomizer"
	"github.com/site-builder/worker/runner"
)

var log = logger.CreateLogger("directory_helper")

type locatorCallback func(locator.Locator)

type DirectoryHelper interface {
	WithTemporaryDirectory(locatorCallback)
}

type directoryHelper struct {
	runner runner.Runner
}

func NewDirectoryHelper(runner runner.Runner) DirectoryHelper {
	return &directoryHelper{runner: runner}
}

func (directoryHelper *directoryHelper) createTempDirectory() locator.Locator {
	randomizer := randomizer.CreateRandomizer()
	directory := locator.NewTempDirectoryLocator(randomizer)

	log.Info("Creating temporary directory %s", directory.Location())

	if err := directoryHelper.runner.Run("mkdir", "-p", directory.Location()); err != nil {
		log.Error("Error creating directory %s", err)
	}

	return directory
}

func (directoryHelper *directoryHelper) removeTempDirectory(directory locator.Locator) {
	log.Info("Removing temporary directory %s", directory.Location())

	if err := directoryHelper.runner.Run("rm", "-rf", directory.Location()); err != nil {
		log.Error("Error removing directory %s", err)
	}
}

func (directoryHelper *directoryHelper) WithTemporaryDirectory(callback locatorCallback) {
	temporaryDirectory := directoryHelper.createTempDirectory()

	callback(temporaryDirectory)

	directoryHelper.removeTempDirectory(temporaryDirectory)
}
