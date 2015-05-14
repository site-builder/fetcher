package locator

import (
	"github.com/site-builder/fetcher/randomizer"

	"fmt"
)

type Locator interface {
	Kind() string
	Location() string
}

type locator struct {
	kind     string
	location string
}

func NewFileLocator(location string) Locator {
	return &locator{location: location, kind: "file"}
}

func NewGitLocator(location string) Locator {
	return &locator{location: location, kind: "git"}
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
