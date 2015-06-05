package builder

import (
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/logger"
	"github.com/site-builder/worker/runner"
)

var log = logger.CreateLogger("builder")

type Builder interface {
	Build(source locator.Locator, destination locator.Locator)
}

type builder struct {
	runner runner.Runner
}

func NewBuilder(runner runner.Runner) Builder {
	return &builder{runner: runner}
}

func (builder *builder) Build(source locator.Locator, destination locator.Locator) {
	log.Info("Building source at %s to %s", source.Location(), destination.Location())

	if err := builder.runner.Run("jekyll", "build", "--source", source.Location(), "--destination", destination.Location()); err != nil {
		log.Error("Error building source %s", err)
	}
}
