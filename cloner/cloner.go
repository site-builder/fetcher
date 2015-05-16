package cloner

import (
	"github.com/site-builder/fetcher/locator"
	"github.com/site-builder/fetcher/logger"
	"github.com/site-builder/fetcher/runner"
)


func Clone(runner runner.Runner, source locator.Locator, destination locator.Locator, log logger.Logger) {
	log.Info("Cloning source from %s to %s", source.Location(), destination.Location())

	if err := runner.Run("git", "clone", "--depth=1", source.Location(), destination.Location()); err != nil {
		log.Error("Error cloning repo %s", err)
	}
}
