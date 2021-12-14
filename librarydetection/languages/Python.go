package languages

import (
	"regexp"

	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/v2/librarydetection"
)

// NewPythonScriptAnalyzer constructor
func NewPythonScriptAnalyzer() librarydetection.Analyzer {
	return &pythonScriptAnalyzer{}
}

type pythonScriptAnalyzer struct{}

func (a *pythonScriptAnalyzer) ExtractLibraries(contents string) ([]string, error) {
	fromRegex, err := regexp.Compile(`from (.+) import`)
	if err != nil {
		return nil, err
	}
	importRegex, err := regexp.Compile(`import ([a-zA-Z0-9_-]+)(?:\s| as)`)
	if err != nil {
		return nil, err
	}

	return executeRegexes(contents, []*regexp.Regexp{fromRegex, importRegex}), nil
}
