package locator

import (
	"github.com/site-builder/fetcher/randomizer"

	"fmt"
)

type Locator interface {
	Kind() string
	Location() string
	Metadata() map[string]string
}

type locator struct {
	kind     string
	location string
	metadata map[string]string
}

func NewFileLocator(location string) Locator {
	return &locator{location: location, kind: "file"}
}

func NewGitLocator(location string) Locator {
	return &locator{location: location, kind: "git", metadata: map[string]string{"branch": "master"}}
}

func NewGitLocatorWithBranch(location string, branch string) Locator {
	return &locator{location: location, kind: "git", metadata: map[string]string{"branch": branch}}
}

func generateTmpDirectoryName(randomizer randomizer.Randomizer) string {
	return fmt.Sprintf("/tmp/%s", randomizer.GenerateUUID())
}

func NewTempDirectoryLocator(randomizer randomizer.Randomizer) Locator {
	return &locator{location: generateTmpDirectoryName(randomizer), kind: "directory"}
}

func (repository *locator) Kind() string {
	return repository.kind
}

func (repository *locator) Location() string {
	return repository.location
}

func (repository *locator) Metadata() map[string]string {
	return repository.metadata
}
