package pkg

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Parse(ctx context.Context, reader *bufio.Reader) error {
	parser := gofeed.NewParser()
	feed, err := parser.Parse(reader)
	if err != nil {
		return err
	}

	for _, item := range feed.Items {
		link := item.Link
		log.Debugf("Link: %s", link)
		response, err := http.Get(fmt.Sprintf("%s.json", link))
		if err != nil {
			return err
		}
		log.Trace(response)

	}

	return nil
}
