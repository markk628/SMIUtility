package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/jinzhu/now"
)

// StockMarketIndex is ...
type StockMarketIndex struct {
	Name  	 string    `json:"name"`
	Index 	 string    `json:"index"`
	DateTime time.Time `json:"date-time"`
}

// CreateJSONFile logs stock info in a json file
func CreateJSONFile(filename string, smis []StockMarketIndex) {
	data, _ := json.MarshalIndent(smis, "", "	")

	if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// ScrapeSMI scrapes stock market indexes' data
func ScrapeSMI(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	time.Now()

	indexes := make([]StockMarketIndex, 0, 3)

	// On every a element which has a.ticker__item.positive attribute call callback
	c.OnHTML("body > div.container.container--zone > div.region.region--primary > div.component.component--module.tickers-bar > div.column.column--full > div.element.element--ticker > div.content-wrapper > div.list.list--tickers > a.ticker__item.positive", func(e *colly.HTMLElement) {

		smiName := e.ChildText("span.label")
		smiPercent := e.ChildText("bg-quote.value")
		dt := now.BeginningOfMinute()
		fmt.Printf("Stock Market Index: %s, Percent Change: %s, Date & Time: %s\n", smiName, smiPercent, dt)
		smi := StockMarketIndex {
			Name:  smiName,
			Index: smiPercent,
			DateTime: dt,
		}
		indexes = append(indexes, smi)

		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(smi)

		w.Header().Set("Content-Type", "application/json")
		w.Write(bf.Bytes())
	})

	// On every a element which has a.ticker__item.negative attribute call callback
	c.OnHTML("body > div.container.container--zone > div.region.region--primary > div.component.component--module.tickers-bar > div.column.column--full > div.element.element--ticker > div.content-wrapper > div.list.list--tickers > a.ticker__item.negative", func(e *colly.HTMLElement) {
	
		smiName := e.ChildText("span.label")
		smiPercent := e.ChildText("bg-quote.value")
		dt := now.BeginningOfMinute()
		fmt.Printf("Stock Market Index: %s, Percent Change: %s, Date & Time: %s\n", smiName, smiPercent, dt)
		smi := StockMarketIndex {
			Name:  smiName,
			Index: smiPercent,
			DateTime: dt,
		}
		indexes = append(indexes, smi)

		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(smi)

		w.Header().Set("Content-Type", "application/json")
		w.Write(bf.Bytes())
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		CreateJSONFile("indexes.json", indexes)
	})

	c.Visit("https://www.marketwatch.com/markets?mod=top_nav")
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	host := "0.0.0.0:8888"
	http.HandleFunc("/", ScrapeSMI)
	fmt.Printf("Localhost: http://%s\n", host)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		return
	}
}
