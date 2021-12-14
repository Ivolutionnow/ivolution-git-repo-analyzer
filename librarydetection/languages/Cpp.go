package languages

import (
	"regexp"

	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/v2/librarydetection"
)

// NewCAnalyzer constructor
func NewCppAnalyzer() librarydetection.Analyzer {
	return &cppAnalyzer{}
}

type cppAnalyzer struct{}

func (a *cppAnalyzer) ExtractLibraries(contents string) ([]string, error) {
	regex, err := regexp.Compile(`(?i)#include\s?[<"]([/a-zA-Z0-9.-]+)[">]`)
	if err != nil {
		return nil, err
	}

	return executeRegexes(contents, []*regexp.Regexp{regex}), nil
}
