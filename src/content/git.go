package content

import (
	"os/exec"
	"fmt"
	"os"
)

type GitClient struct {
	Path   string
	GitUrl string
}

func (gc GitClient) Clone() error {
	cmd := exec.Command("/usr/bin/git", "clone", gc.GitUrl, gc.Path)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (gc GitClient) Fetch() {

}

func (gc GitClient) Pull() {

}

func (gc GitClient) CommitHash() (string, error) {
	cmd := exec.Command("/usr/bin/git", "rev-parse", "HEAD")
	cmd.Dir = gc.Path
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
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
