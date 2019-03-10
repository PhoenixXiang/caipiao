package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"

	. "github.com/PhoenixXiang/caipiao/util"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.55128.cn/zs/12_96.htm?startTerm=2007001&endTerm=2007002")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	ts := make([]Ticket, 0, 5)
	doc.Find("#chartData tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		brl := s.Find(".brl").Text()
		red := [5]int{}
		s.Find("[class^='chartball_red']").Each(func(i int, s *goquery.Selection) {
			red[i], _ = strconv.Atoi(s.Text())
		})
		blue := [2]int{}
		s.Find("[class^='chartball_blue']").Each(func(i int, s *goquery.Selection) {
			blue[i], _ = strconv.Atoi(s.Text())
		})
		
		// title := s.Find("i").Text()
		fmt.Printf("%s期 - %v - %v\n", brl, red, blue)
		ts = append(ts, GetTicket(brl, red[:], blue[:]))
	})
	Marshal(ts)
}

func getNet() {
	res, err := http.Get("https://www.55128.cn/zs/12_96.htm?startTerm=2007001&endTerm=2019026")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func main() {
	// getNet()
	ExampleScrape()
}
