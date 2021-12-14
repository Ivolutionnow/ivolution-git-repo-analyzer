package extractor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Ivolutionnow/ivolution-git-repo-analyzer/v2/extractor"
)

var _ = Describe("GetRepoName", func() {

	Context("RepoExtractor headless", func() {
		It("should get the repo name with the owner name", func() {
			re := extractor.RepoExtractor{
				Headless: true,
			}
			Expect(re.GetRepoName("git@github.com:Ivolutionnow/ivolution-git-repo-analyzer.git")).To(Equal("Ivolutionnow/ivolution-git-repo-analyzer"))
			Expect(re.GetRepoName("https://github.com/Ivolutionnow/ivolution-git-repo-analyzer.git")).To(Equal("Ivolutionnow/ivolution-git-repo-analyzer"))
			Expect(re.GetRepoName("https://github.com/chelovekula/second-project.git")).To(Equal("chelovekula/second-project"))
			Expect(re.GetRepoName("ssh://user@host:port/group/repoName.git")).To(Equal("group/repoName"))
		})
	})

	Context("RepoExtractor interactive", func() {
		It("should get the repo name without the owner name", func() {
			re := extractor.RepoExtractor{
				Headless: false,
				RepoPath: "/some/path/Ivolutionnow/ivolution-git-repo-analyzer",
			}
			Expect(re.GetRepoName("git@github.com:Ivolutionnow/ivolution-git-repo-analyzer.git")).To(Equal("ivolution-git-repo-analyzer"))
			Expect(re.GetRepoName("https://github.com/Ivolutionnow/ivolution-git-repo-analyzer.git")).To(Equal("ivolution-git-repo-analyzer"))
			Expect(re.GetRepoName("")).To(Equal("ivolution-git-repo-analyzer"))
			Expect(re.GetRepoName("ssh://user@host:port/group/repoName.git")).To(Equal("repoName"))
		})
	})
})
