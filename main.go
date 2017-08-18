package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adlio/darksky"
)

func CompareDate(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func main() {
	latPtr := flag.String("latitude", "42.8864", "latitude")
	lngPtr := flag.String("longitude", "-78.8784", "longitude")
	epochPtr := flag.Bool("epoch", true, "Return SunsetTime in epoch")

	flag.Parse()

	apiKey := os.Getenv("DS_APIKEY")
	if apiKey == "" {
		log.Fatalln("Need a Dark SKy API Key")
	}

	client := darksky.NewClient(apiKey)
	f, err := client.GetForecast(*latPtr, *lngPtr, darksky.Defaults)
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()
	for _, elm := range f.Daily.Data {
		elmDate := time.Unix(int64(elm.SunsetTime), 0)
		if CompareDate(now, elmDate) {
			if *epochPtr {
				fmt.Printf("%v\n", elmDate.Unix())
			} else {
				fmt.Printf("%v\n", elmDate)
			}
			return
		}
	}
}