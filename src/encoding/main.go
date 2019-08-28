package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rss struct {
	XMLName  xml.Name  `xml:"rss"`
	Channels []Channel `xml:"channel"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Language    string   `xml:"language"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Categories  []string `xml:"category"`
	// Creator     string   `xml:"dc:creator,cdata"`
}

func main() {
	resp, err := http.Get("https://es.gizmodo.com/rss")
	if err != nil {
		fmt.Println("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status error: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body error: %v", err)
	}

	var rss Rss
	xml.Unmarshal(body, &rss)

	for _, val := range rss.Channels[0].Items {
		fmt.Println(val.Title)
	}
}
