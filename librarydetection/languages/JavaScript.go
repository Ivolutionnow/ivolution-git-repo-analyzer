package languages

import (
	"regexp"

	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/v2/librarydetection"
)

// NewJavaScriptAnalyzer constructor
func NewJavaScriptAnalyzer() librarydetection.Analyzer {
	return &javaScriptAnalyzer{}
}

type javaScriptAnalyzer struct{}

func (a *javaScriptAnalyzer) ExtractLibraries(contents string) ([]string, error) {

	require, err := regexp.Compile(`require\(["\'](.+)["\']\);?\s`)
	if err != nil {
		return nil, err
	}

	importRegex, err := regexp.Compile(`import\s*(?:.+ from)?\s?\(?[\'"](.+)[\'"]\)?;?\s`)
	if err != nil {
		return nil, err
	}

	return executeRegexes(contents, []*regexp.Regexp{require, importRegex}), nil
}
