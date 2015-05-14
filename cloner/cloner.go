package cloner

import (
	"github.com/site-builder/fetcher/commander"
	"github.com/site-builder/fetcher/locator"
	"github.com/site-builder/fetcher/logger"
)

func Clone(source locator.Locator, destination locator.Locator, log logger.Logger) {
	command := commander.CreateCommand("git", "clone", "--depth=1", source.Location(), destination.Location())

	log.Info("Cloning source from %s to %s", source.Location(), destination.Location())

	if err := command.Run(); err != nil {
		log.Error("Error cloning repo %s", err)
	}
}
