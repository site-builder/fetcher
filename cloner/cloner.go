package cloner

import (
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/logger"
	"github.com/site-builder/worker/runner"
)

var log = logger.CreateLogger("cloner")

func Clone(runner runner.Runner, source locator.Locator, destination locator.Locator) {
	log.Info("Cloning source from %s to %s", source.Location(), destination.Location())
	branch := source.Metadata()["branch"]

	if err := runner.Run("git", "clone", "-b", branch, "--single-branch", "--depth=1", source.Location(), destination.Location()); err != nil {
		log.Error("Error cloning repo %s", err)

		if message := err.Error(); message == "exit status 128" {
			log.Info("Does the repo you're trying to clone exist?")
		}
	}
}
