package content

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/otsimo/api/apipb"
)

type ContentConfig struct {
	Folders      []string
	Template     string
	PublicDirs   map[string]string
	AssetVersion int32
}

type PublicDirEntry struct {
	//Path is shown on http url
	Path string
	//Dir is real path where folder is located
	Dir string
}

type ContentManager struct {
	Git           *GitClient
	CurrentCommit string
	GitPublicDirs []PublicDirEntry
	publicDir     string
	host          string
	contents      []*apipb.Content
	tempContents  []*apipb.Content
	assetVersion  int32
}

func NewContentManager(config *Config) *ContentManager {
	git := NewGitClient(config.GitFolder, config.GitUrl)

	return &ContentManager{
		Git:           git,
		publicDir:     config.PublicDir,
		GitPublicDirs: []PublicDirEntry{},
		host:          config.Host,
		contents:      []*apipb.Content{},
		tempContents:  []*apipb.Content{},
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
	//cm.ClearPublicDir()

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

func (cm *ContentManager) readDirectory(dir string, tpl *template.Template) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".md" {
			content, err := parseMarkdownFile(path, cm.publicDir, tpl)
			if err != nil {
				log.Errorf("failed to parse markdown file '%s' err=%v", path, err)
				return nil
			}
			if content != nil {
				cm.AddContent(content)
			}
		}
		return nil
	})
}

func contentHtmlFilename(content *apipb.Content) string {
	return content.Slug + "." + content.Language + ".html"
}

func (cm *ContentManager) ClearPublicDir() {
	if ex, _ := pathExists(cm.publicDir); ex {
		os.RemoveAll(cm.publicDir)
	}
	os.Mkdir(cm.publicDir, os.ModePerm)
}

func (cm *ContentManager) ReadContent() error {
	configPath := filepath.Join(cm.Git.Path, "config.json")
	config, err := NewContentConfig(configPath)
	if err != nil {
		return err
	}

	for k, v := range config.PublicDirs {
		cm.GitPublicDirs = append(cm.GitPublicDirs, PublicDirEntry{
			Dir:  filepath.Join(cm.Git.Path, v),
			Path: k,
		})
	}

	tpl, err := ioutil.ReadFile(filepath.Join(cm.Git.Path, config.Template))
	if err != nil {
		log.Errorf("failed to read template")
		return err
	}
	templ, err := template.New("webpage").Parse(string(tpl))
	if err != nil {
		log.Errorf("failed to parse template")
		return err
	}
	for _, v := range config.Folders {
		dp := path.Join(cm.Git.Path, v)
		cm.readDirectory(dp, templ)
	}
	cm.assetVersion = config.AssetVersion
	cm.contents = cm.tempContents
	cm.tempContents = []*apipb.Content{}
	return nil
}

func (cm *ContentManager) AddContent(content *apipb.Content) error {
	content.Url = cm.host + WikiEndpoint + "/" + contentHtmlFilename(content)

	cm.tempContents = append(cm.tempContents, content)

	return nil
}

func (cm *ContentManager) Update(commit string) error {
	err := cm.Git.Checkout(commit)
	if err != nil {
		log.Errorf("content.go: failed to checkout repository")
		return err
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
