package cloner

import (
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/logger"
	"github.com/site-builder/worker/runner"
)

var log = logger.CreateLogger("cloner")

func Clone(runner runner.Runner, source locator.Locator, destination locator.Locator) {
	branch := source.Metadata()["branch"]
	log.Info("Cloning source from %s@%s to %s", source.Location(), branch, destination.Location())

	if err := runner.Run("git", "clone", "-b", branch, "--single-branch", "--depth=1", source.Location(), destination.Location()); err != nil {
		log.Error("Error cloning repo %s", err)

		if message := err.Error(); message == "exit status 128" {
			log.Info("Does the repo you're trying to clone exist?")
		}
	}
}
