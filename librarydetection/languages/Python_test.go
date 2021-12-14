package languages_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"

	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/v2/librarydetection/languages"
)

var _ = Describe("PythonLibraryDetection", func() {
	fixture, err := ioutil.ReadFile("./fixtures/python.fixture")
	if err != nil {
		panic(err)
	}

	expectedLibraries := []string{
		"lib1.lib2",
		"lib3",
		"lib4",
	}

	analyzer := languages.NewPythonScriptAnalyzer()

	Describe("Extract Python Libraries", func() {
		It("Should be able to extract libraries", func() {
			libs, err := analyzer.ExtractLibraries(string(fixture))
			if err != nil {
				panic(err)
			}
			assertSameUnordered(libs, expectedLibraries)
		})
	})
})
