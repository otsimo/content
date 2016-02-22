package content

import (
	"encoding/json"
	"os"
	"text/template"

	"path/filepath"

	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/otsimo/api/apipb"
	"github.com/russross/blackfriday"
)

func parseMarkdownFile(filep string, outdir string, tpl *template.Template) (*apipb.Content, error) {
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
		time.RFC3339,
		fmt.Sprintf("%s:00+00:00", m["writtenAt"]))

	if e != nil {
		return nil, fmt.Errorf("invalid time format:%v", e)
	}

	content.Date = t1.Unix()
	content.Markdown = page.Content()
	out := blackfriday.MarkdownCommon(content.Markdown)

	data := struct {
		Title   string
		Content string
	}{
		Title:   content.Title,
		Content: string(out),
	}
	outfilepath := filepath.Join(outdir, contentHtmlFilename(content))
	ofile, err := os.Create(outfilepath)
	if err != nil {
		log.Errorf("failed create to output file, path=%s, err=%v", outfilepath, err)
		return nil, err
	}

	err = tpl.Execute(ofile, data)
	if err != nil {
		log.Errorf("failed to execute template")
	}

	return &content, nil
}