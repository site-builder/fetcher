package deployer

import (
	"github.com/site-builder/worker/locator"
	"github.com/site-builder/worker/logger"
	"github.com/site-builder/worker/runner"
)

var log = logger.CreateLogger("deployer")

type Deployer interface {
	Deploy(source locator.Locator, destination locator.Locator)
}

type deployer struct {
	runner runner.Runner
}

func NewDeployer(runner runner.Runner) Deployer {
	return &deployer{runner: runner}
}

func (deployer *deployer) Deploy(source locator.Locator, destination locator.Locator) {
	branch := destination.Metadata()["branch"]

	log.Info("Deploying from %s to %s@%s", source.Location(), destination.Location(), branch)

	deployer.runner.Run("git", "-C", source.Location(), "init")
	deployer.runner.Run("git", "-C", source.Location(), "add", ".")
	deployer.runner.Run("git", "-C", source.Location(), "commit", "-m", "built using site-builder")
	deployer.runner.Run("git", "-C", source.Location(), "remote", "add", "origin", destination.Location())

	if err := deployer.runner.Run("git", "-C", source.Location(), "push", "origin", "master:"+branch, "--force"); err != nil {
		log.Error("Error 5 %s", err)
	}
}
