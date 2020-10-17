package pkg

import (
	"bufio"
	"bytes"
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
)

func Parse(reader *bufio.Reader) error {
	parser := gofeed.NewParser()
	feed, err := parser.Parse(reader)
	if err != nil {
		return err
	}

	article := feed.Items[0]
	log.Tracef("Article: %+v", article)
	content := article.Content

	//stringReader := strings.NewReader(content)
	//
	//formatted, err := readability.FromReader(stringReader, "http://abc.com")
	//if err != nil {
	//	return err
	//}
	//
	//log.Tracef("Article content: %s", formatted.Content)
	//
	//d := []byte(formatted.Content)

	d := []byte(content)

	command := exec.Command("pandoc", "-f", "html", "-t", "epub")
	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdin = bytes.NewBuffer(d)
	command.Stdout = &out
	command.Stderr = &stderr
	err = command.Run()
	if err != nil {
		log.Debugf("errors: %s", string(stderr.Bytes()))
		return err
	}
	file, err := ioutil.TempFile(".", "*.epub")
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(out.Bytes())
	if err != nil {
		return err
	}

	return nil
}
