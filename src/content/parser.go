package content

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	log "github.com/Sirupsen/logrus"
	apipb "github.com/otsimo/otsimopb"
	"github.com/russross/blackfriday"
)

func parseMarkdownFile(filep string, cm *ContentManager, tpl *template.Template) (*apipb.Content, error) {
	file, err := os.Open(filep)
	if err != nil {
		log.Errorf("failed to open file %s error:%v", filep, err)
		return nil, err
	}

	page, err := ReadFrom(file)
	if err != nil {
		log.Errorf("failed to read page %v", err)
		return nil, err
	}
	m, err := page.Metadata()
	if err != nil {
		log.Errorf("failed to get metadata %v", err)
		return nil, err
	}
	d, err := json.Marshal(m)
	if err != nil {
		log.Errorf("failed to marshall metadata, error=%v", err)
		return nil, err
	}
	content := apipb.Content{}
	if err := json.Unmarshal(d, &content); err != nil {
		log.Errorf("failed marshall content metadata, error=%v", err)
		return nil, err
	}

	t1, e := time.Parse(
		"2006-01-02",
		content.WrittenAt)

	if e != nil {
		log.Errorf("failed to parse date, error=%v", e)
		return nil, fmt.Errorf("invalid time format:%v", e)
	}

	content.Date = t1.Unix()
	content.Markdown = page.Content()
	out := blackfriday.MarkdownCommon(content.Markdown)

	data := struct {
		Title   string
		Content string
		Params  map[string]string
		Author  string
	}{
		Title:   content.Title,
		Content: string(out),
		Params:  content.Params,
		Author:  content.Author,
	}
	outfilepath := filepath.Join(cm.publicDir, cm.contentHtmlFilename(&content))
	ofile, err := os.Create(outfilepath)
	if err != nil {
		log.Errorf("failed create to output file, path=%s, err=%v", outfilepath, err)
		return nil, err
	}

	err = tpl.Execute(ofile, data)
	if err != nil {
		log.Errorf("failed to execute template: %v", err)
	}

	return &content, nil
}
