package content

import (
	log "github.com/Sirupsen/logrus"
	"errors"
)

type ContentManager struct {
	Git           *GitClient
	CurrentCommit string
}

func NewContentManager(config *Config) *ContentManager {
	git := NewGitClient(config.GitFolder, config.GitUrl)

	return &ContentManager{
		Git: git,
	}
}

func (cm *ContentManager) Init() error {
	if cm.Git.HasGitRepository(cm.Git.Path) {
		remote, err := cm.Git.RemoteUrl(cm.Git.Path)
		if err != nil {
			log.Errorf("failed to get remote url error:%v", err)
			return err
		}
		if remote != cm.Git.GitUrl {
			log.Errorf("git repository at given path does not match, want='%s' got='%s'", cm.Git.GitUrl, remote)
			return errors.New("git repository at given path does not match")
		}
		err = cm.Git.FetchAndPull()
		if err != nil {
			log.Errorf("content.go: failed to pull repository")
			return err
		}
	}else {
		err := cm.Git.Clone()
		if err != nil {
			log.Errorf("content.go: failed to clone repository")
			return err
		}
	}

	ch, err := cm.Git.CommitHash()
	if err == nil {
		cm.CurrentCommit = ch
	}

	log.Infof("content.go: Git.CommitHash '%s', '%v'", ch, err)

	return err
}
