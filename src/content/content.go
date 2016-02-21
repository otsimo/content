package content

import (
	"errors"

	"encoding/json"
	"os"
	"path"
	"path/filepath"

	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type ContentConfig struct {
	Folders []string
}

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
	} else {
		err := cm.Git.Clone()
		if err != nil {
			log.Errorf("content.go: failed to clone repository")
			return err
		}
	}

	ch, err := cm.Git.CommitHash()
	if err == nil {
		cm.CurrentCommit = ch
	} else {
		return err
	}

	log.Infof("content.go: Git.CommitHash '%s', '%v'", ch, err)
	return cm.ReadContent()
}

func NewContentConfig(configPath string) (*ContentConfig, error) {
	config := ContentConfig{}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Errorf("failed to open file, error:%v", err)
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Errorf("failed to decode json file, error:%v", err)
		return nil, err
	}
	return &config, nil
}

func readDirectory(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".md" {
			fmt.Printf("%s %v\n", path, err)
		}
		return nil
	})
}

func (cm *ContentManager) ReadContent() error {
	configPath := filepath.Join(cm.Git.Path, "config.json")
	config, err := NewContentConfig(configPath)

	if err != nil {
		return err
	}

	for _, v := range config.Folders {
		dp := path.Join(cm.Git.Path, v)
		readDirectory(dp)
	}
	return nil
}
