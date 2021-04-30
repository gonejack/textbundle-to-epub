package cmd

import (
	"bytes"
	"errors"
	"fmt"

	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gomarkdown/markdown"
	"github.com/gonejack/go-epub"
)

type TextBundleToEpub struct {
	DefaultCover []byte

	Cover   string
	Title   string
	Author  string
	Verbose bool

	book *epub.Epub
}

func (t *TextBundleToEpub) Run(textBundles []string, output string) (err error) {
	if len(textBundles) == 0 {
		return errors.New("no textbundle given")
	}

	t.book = epub.NewEpub(t.Title)
	{
		t.setAuthor()
		t.setDesc()
		err = t.setCover()
		if err != nil {
			return
		}
	}

	for _, textbundle := range textBundles {
		err = t.appendTextBundle(textbundle)
		if err != nil {
			err = fmt.Errorf("parse %s failed: %s", textbundle, err)
			return
		}
	}

	err = t.book.Write(output)
	if err != nil {
		return fmt.Errorf("cannot write output epub: %s", err)
	}

	return
}
func (t *TextBundleToEpub) appendTextBundle(textbundle string) (err error) {
	basedir := textbundle

	mdf, err := os.Open(filepath.Join(basedir, "text.markdown"))
	if err != nil {
		return
	}

	md, err := ioutil.ReadAll(mdf)
	if err != nil {
		return
	}

	htm := t.md2Html(md)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htm))
	if err != nil {
		return
	}
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		ref := filepath.Join(basedir, src)

		// be compatible with bear
		if _, err := os.Stat(ref); errors.Is(err, os.ErrNotExist) {
			src, _ = url.PathUnescape(src)
			ref = filepath.Join(basedir, src)
		}

		internalRef, _ := t.book.AddImage(ref, "")

		selection.SetAttr("src", internalRef)
	})

	content, err := doc.Html()
	if err != nil {
		return
	}

	title := filepath.Base(basedir)
	title = strings.TrimSuffix(title, filepath.Ext(title))
	_, err = t.book.AddSection(content, title, "", "")

	return
}
func (t *TextBundleToEpub) md2Html(md []byte) (html []byte) {
	return markdown.ToHTML(md, nil, nil)
}

func (t *TextBundleToEpub) setAuthor() {
	t.book.SetAuthor(t.Author)
}
func (t *TextBundleToEpub) setDesc() {
	t.book.SetDescription(fmt.Sprintf("Epub generated at %s with github.com/gonejack/textbundle-to-epub", time.Now().Format("2006-01-02")))
}
func (t *TextBundleToEpub) setCover() (err error) {
	if t.Cover == "" {
		temp, err := os.CreateTemp("", "textbundle-to-epub")
		if err != nil {
			return fmt.Errorf("cannot create tempfile: %s", err)
		}
		_, err = temp.Write(t.DefaultCover)
		if err != nil {
			return fmt.Errorf("cannot write tempfile: %s", err)
		}
		_ = temp.Close()

		t.Cover = temp.Name()
	}

	fmime, err := mimetype.DetectFile(t.Cover)
	if err != nil {
		return fmt.Errorf("cannot detect cover mime type %s", err)
	}

	coverRef, err := t.book.AddImage(t.Cover, "epub-cover"+fmime.Extension())
	if err != nil {
		return fmt.Errorf("cannot add cover %s", err)
	}

	t.book.SetCover(coverRef, "")

	return
}
