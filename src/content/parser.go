package content

import (
	"encoding/json"
	"os"
	"text/template"

	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/otsimo/api/apipb"
	"github.com/russross/blackfriday"
)

func parseMarkdownFile(filep string, outdir string, tpl *template.Template) (*Page, *apipb.Content, error) {
	file, err := os.Open(filep)
	if err != nil {
		log.Errorf("failed to open file %s error:%v", filep, err)
		return nil, nil, err
	}

	page, err := ReadFrom(file)
	if err != nil {
		log.Errorf("failed to read page %v", err)
		return nil, nil, err
	}
	m, err := page.Metadata()
	if err != nil {
		log.Errorf("failed to get metadata %v", err)
		return nil, nil, err
	}
	d, err := json.Marshal(m)
	if err != nil {
		log.Errorf("failed to marshall metadata, error=%v", err)
		return nil, nil, err
	}
	content := apipb.Content{}
	if err := json.Unmarshal(d, &content); err != nil {
		log.Errorf("failed marshall content metadata, error=%v", err)
		return nil, nil, err
	}
	content.Markdown = page.Content()
	out := blackfriday.MarkdownCommon(content.Markdown)

	data := struct {
		Title   string
		Content string
	}{
		Title:   content.Title,
		Content: string(out),
	}
	log.Infoln("markdown->html", data.Content)
	outfilepath := filepath.Join(outdir, content.Slug+"-"+content.Language+".html")
	ofile, err := os.Create(outfilepath)
	if err != nil {
		log.Errorf("failed create to output file, path=%s, err=%v", outfilepath, err)
		return nil, nil, err
	}

	err = tpl.Execute(ofile, data)
	if err != nil {
		log.Errorf("failed to execute template")
	}

	return &page, &content, nil
}
