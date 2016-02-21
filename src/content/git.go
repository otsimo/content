package content

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
)

type GitClient struct {
	Path   string
	GitUrl string
}

func (gc GitClient) Clone() error {
	cmd := exec.Command("/usr/bin/git", "clone", gc.GitUrl, gc.Path)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (gc GitClient) Fetch() {

}

func (gc GitClient) Pull() {

}

func (gc GitClient) CommitHash() (string, error) {
	cmd := exec.Command("/usr/bin/git", "rev-parse", "HEAD")
	cmd.Dir = gc.Path
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := string(b)
	if strings.HasSuffix(str, "\n") {
		str = str[:(len(str) - 1)]
	}
	return str, nil
}

func (gc GitClient) HasCommitHash(sha1 string) bool {
	cmd := exec.Command("/usr/bin/git", "log", sha1, "-n", "1")
	cmd.Dir = gc.Path
	err := cmd.Run()
	return err == nil
}

//git rev-parse HEAD
//git log 491652ab004e4b3fee6976f48838c2e78d580294 -n 1
func NewGitClient(folder, url string) *GitClient {
	return &GitClient{Path: folder, GitUrl: url}
}
