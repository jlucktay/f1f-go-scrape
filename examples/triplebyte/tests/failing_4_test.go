package main

import (
	"net/url"
	"testing"

	"github.com/jlucktay/f1-go-scrape/examples/triplebyte/crawler"
)

func TestChallenge4(t *testing.T) {
	c := crawler.Crawler{
		Threads: 5,
		Log:     crawler.Verbose(),
	}
	graph, err := c.Crawl("http://triplebyte.github.io/web-crawler-test-site/test4/", "")
	if err != nil {
		t.Fatal("Broken test, can't run crawl")
	}

	u, err := url.Parse("https://triplebyte.github.io/web-crawler-test-site/test4/page3")
	if err != nil {
		t.Fatal("Broken test, can't parse URL")
	}

	if _, ok := graph.Nodes[*u]; !ok {
		t.Errorf("Not found: %v", u)
	}
}
