package locator_test

import (
	. "github.com/site-builder/fetcher/locator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type mockRandomizer struct {
	uuid string
}

func (mockRandomizer *mockRandomizer) GenerateUUID() string {
	return mockRandomizer.uuid
}

var _ = Describe("Locator", func() {
	Context("#NewGitLocator", func() {
		location := "git@github.com:site-builder/fetcher.git"

		var locator Locator

		It("sets the Kind to 'git'", func() {
			locator = NewGitLocator(location)
			Expect(locator.Kind()).To(Equal("git"))
		})

		It("sets the Location", func() {
			locator = NewGitLocator(location)
			Expect(locator.Location()).To(Equal(location))
		})

		It("sets the branch in the Metadata to master", func() {
			locator = NewGitLocator(location)
			Expect(locator.Metadata()).To(Equal(map[string]string{"branch": "master"}))
		})
	})

	Context("#NewGitLocatorWithBranch", func() {
		location := "git@github.com:site-builder/fetcher.git"

		var locator Locator

		It("sets the Kind to 'git'", func() {
			locator = NewGitLocatorWithBranch(location, "other")
			Expect(locator.Kind()).To(Equal("git"))
		})

		It("sets the Location", func() {
			locator = NewGitLocatorWithBranch(location, "other")
			Expect(locator.Location()).To(Equal(location))
		})

		It("sets the branch in the Metadata to master", func() {
			locator = NewGitLocatorWithBranch(location, "other")
			Expect(locator.Metadata()).To(Equal(map[string]string{"branch": "other"}))
		})
	})

	Context("#NewFileLocator", func() {
		location := "file:///tmp"

		var locator Locator

		It("sets the Kind to 'git'", func() {
			locator = NewFileLocator(location)
			Expect(locator.Kind()).To(Equal("file"))
		})

		It("sets the Location", func() {
			locator = NewFileLocator(location)
			Expect(locator.Location()).To(Equal(location))
		})
	})

	Context("#NewTempDirectoryLocator", func() {
		It("sets the Kind to 'directory'", func() {
			randomizer := &mockRandomizer{uuid: "my-mock-uuid"}
			locator := NewTempDirectoryLocator(randomizer)
			Expect(locator.Kind()).To(Equal("directory"))
		})

		It("sets the Location to /tmp/my-mock-uuid", func() {
			randomizer := &mockRandomizer{uuid: "my-mock-uuid"}
			locator := NewTempDirectoryLocator(randomizer)
			Expect(locator.Location()).To(Equal("/tmp/my-mock-uuid"))
		})
	})
})
