package content

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
)

const (
	gitPath = "/usr/bin/git"
)

type GitClient struct {
	Path   string
	GitUrl string
}

func (gc GitClient) Clone() error {
	log.Debugf("git.go: starting to clone %s to %s", gc.GitUrl, gc.Path)
	cmd := exec.Command(gitPath, "clone", gc.GitUrl, gc.Path)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (gc GitClient) FetchAndPull() error {
	cmd := exec.Command(gitPath, "fetch")
	cmd.Dir = gc.Path
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command(gitPath, "pull")
	cmd.Dir = gc.Path
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (gc GitClient) CommitHash() (string, error) {
	cmd := exec.Command(gitPath, "rev-parse", "HEAD")
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
	cmd := exec.Command(gitPath, "log", sha1, "-n", "1")
	cmd.Dir = gc.Path
	err := cmd.Run()
	return err == nil
}
func (gc GitClient) Checkout(arg ...string) error {
	log.Debugf("git.go: starting to checkout %s to %s", gc.GitUrl, gc.Path, arg)
	args := append([]string{"checkout"}, arg...)
	cmd := exec.Command(gitPath, args...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (gc GitClient) RemoteUrl(path string) (string, error) {
	cmd := exec.Command(gitPath, "config", "--get", "remote.origin.url")
	cmd.Dir = path
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

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func (gc GitClient) HasGitRepository(dir string) bool {
	gitpath := filepath.Join(dir, ".git")
	ex, _ := pathExists(gitpath)
	return ex
}

//git rev-parse HEAD
//git log 491652ab004e4b3fee6976f48838c2e78d580294 -n 1
func NewGitClient(folder, url string) *GitClient {
	return &GitClient{Path: folder, GitUrl: url}
}
