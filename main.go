package main

import (
	"flag"
	"strings"

	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/autoupdater"
	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/extractor"
)

var (
	version = "v9.9.9" // Version of the file. E.g v0.9.6. This is set during build time.
)

func main() {
	au := autoupdater.NewAutoUpdater(version)
	au.CheckUpdates()

	repoPath := flag.String("repo_path", "", "Path of the repo")
	// Following two flags should be used to disable email prompt
	// Program is going to ask you to choose your emails
	// But if you want, you can provide the emails yourself
	headless := flag.String("headless", "false", "Headless mode is used on Ivolution backend system.")
	obfuscate := flag.String("obfuscate", "true", "Set it to true for debug purposes.")
	outputPath := flag.String("output_path", "./repo_data_v2", "Where to put output file")
	gitPath := flag.String("git_path", "", "Where is git binary?")
	emailString := flag.String("emails", "", "Predefined emails. Example: \"devs@ivolution.io,devs@gmail.com\"")
	seeds := flag.String("seeds", "", "The seed is used to find similar emails. Example: \"devs, devs@ivolution.io\"")
	repoName := flag.String("repo_name", "", "You can overwrite the default repo name. This name will be shown on the profile page.")
	skipLibraries := flag.Bool("skip_libraries", false, "Turns off the library detection in order to reduce the execution time.")
	flag.Parse()

	if repoPath == nil || *repoPath == "" {
		panic("Please provide a path to the repo")
	}

	emails := make([]string, 0)
	if emailString != nil && len(*emailString) > 0 {
		emails = strings.Split(*emailString, ",")
	}

	seed := make([]string, 0)
	if seeds != nil && len(*seeds) > 0 {
		seed = strings.Split(*seeds, ",")
	}

	repoExtractor := extractor.RepoExtractor{
		RepoPath:            *repoPath,
		OutputPath:          *outputPath,
		GitPath:             *gitPath,
		Headless:            *headless == "true",
		Obfuscate:           *obfuscate == "true",
		UserEmails:          emails,
		Seed:                seed,
		ShowProgressBar:     *headless != "true", // Show progress bar only if running in interactive mode
		OverwrittenRepoName: *repoName,
		SkipLibraries:       *skipLibraries,
	}

	err := repoExtractor.Extract()
	if err != nil {
		panic(err)
	}
}
