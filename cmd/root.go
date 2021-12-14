package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type rootConfig struct {
	SkipLibraries *bool
	SkipUpdate    *bool
	Seeds         *[]string
	Emails        *[]string
	GitPath       *string
	OutPutPath    *string
	Obfuscate     *bool
	Headless      *bool
}

var (
	rootCmd = &cobra.Command{
		Use:   "ivolution-git-repo-analyzer",
		Short: "Extract data from a Git repository",
		Long: `Use this command to extract and upload repo data your ivolution profile.
Example usage: ivolution-git-repo-analyzer path --repo_path /path/to/repo`,
	}

	RootConfig rootConfig

	emailString *string
	seedsString *string
	Version     string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootConfig.SkipLibraries = rootCmd.PersistentFlags().Bool("skip_libraries", false, "Turns off the library detection in order to reduce the execution time")
	RootConfig.SkipUpdate = rootCmd.PersistentFlags().Bool("skip_update", false, "If set the auto-update is skipped")
	emailString = rootCmd.PersistentFlags().String("emails", "", "Predefined emails. Example: \"dmitry@ivolution.ai,dmitry@gmail.com\"")
	seedsString = rootCmd.PersistentFlags().String("seeds", "", "The seed is used to find similar emails. Example: \"Ivolutionnow, Ivolutionnow@ivolution.ai\"")
	RootConfig.GitPath = rootCmd.PersistentFlags().String("git_path", "", "where the Git binary is")
	RootConfig.OutPutPath = rootCmd.PersistentFlags().String("output_path", "./artifacts", "Where to put output file. Existing artifacts will be overwritten.")
	RootConfig.Obfuscate = rootCmd.PersistentFlags().Bool("obfuscate", true, "File names and emails won't be hashed. Set it to true for debug purposes.")
	RootConfig.Headless = rootCmd.PersistentFlags().Bool("headless", false, "Headless mode is used on ivolution's backend system.")
}

func initConfig() {
	emails := make([]string, 0)
	if len(*emailString) > 0 {
		emails = strings.Split(*emailString, ",")
	}
	RootConfig.Emails = &emails

	seeds := make([]string, 0)
	if len(*seedsString) > 0 {
		seeds = strings.Split(*seedsString, ",")
	}

	RootConfig.Seeds = &seeds

	// Find git executable if it is not provided
	if *RootConfig.GitPath == "" {
		gitPath, err := exec.LookPath("git")
		if err != nil {
			defaultGitPath := "/usr/bin/git"
			fmt.Printf("Couldn't find git path. Fall back to default (%s). Error: %s.\n", defaultGitPath, err.Error())
			// Try default git path
			*RootConfig.GitPath = defaultGitPath
			return
		}
		gitPath = strings.TrimRight(gitPath, "\r\n")
		gitPath = strings.TrimRight(gitPath, "\n")

		*RootConfig.GitPath = gitPath
	}
}
