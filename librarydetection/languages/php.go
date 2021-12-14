package languages

import (
	"regexp"
	"strings"

	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/v2/librarydetection"
)

// NewCAnalyzer constructor
func NewPHPAnalyzer() librarydetection.Analyzer {
	return &phpAnalyzer{}
}

type phpAnalyzer struct{}

func (a *phpAnalyzer) ExtractLibraries(contents string) ([]string, error) {
	// matches
	// require('lib1');
	// require 'lib2';
	// require("lib3");
	// require "lib4";
	// include('lib5');
	// include 'lib6';
	// include("lib7");
	// include "lib8";
	// require_once('lib9');
	// require_once 'lib10';
	// require_once("lib11");
	// require_once "lib12";
	regex1, err := regexp.Compile(`(?i)(?:require|require_once|include)[( ]{1}['"]([a-zA-Z0-9]+)["'][)]?;`)
	if err != nil {
		return nil, err
	}

	// match all `use` imports
	regex2, err := regexp.Compile(`(?i)use ([a-zA-Z]+\\)[^;]*`)
	if err != nil {
		return nil, err
	}

	ret := executeRegexes(contents, []*regexp.Regexp{regex1, regex2})

	var res []string
	for _, v := range ret {
		if !strings.HasPrefix(v, "App") {
			res = append(res, v)
		}
	}

	return res, nil
}
