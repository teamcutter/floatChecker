package main

import (
	"encoding/json"
	"floatChecker/requests"
	"fmt"
	_ "io"
	"io/ioutil"
	"math"
	_ "math"
	_ "strings"
	"sync"
	"time"
)

const thousand = "https://steamcommunity.com/market/listings/730/AK-47%20%7C%20Slate%20%28Minimal%20Wear%29/render/?query=country=EU&language=english&currency=1"

func main() {
	startTime := time.Now()
	// тут много https://steamcommunity.com/market/listings/730/StatTrak%E2%84%A2%20AK-47%20%7C%20Uncharted%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1
	url := "https://steamcommunity.com/market/listings/730/StatTrak™%20Desert%20Eagle%20%7C%20Directive%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1"
	links := requests.SearchCurrentItem(url)
	fmt.Println("Links: ", len(links))

	var wg sync.WaitGroup

	flCh := make(chan requests.FloatInfo)
	var floatInfoList []requests.FloatInfo

	start := 0

	countOfGoRoutines := int(math.Ceil(float64(len(links)) / 100))
	fmt.Println("Count of goroutines: ", countOfGoRoutines)

	if len(links) > 100 {
		for i := 0; i < countOfGoRoutines; i++ {
			wg.Add(1)
			count := len(links) - start
			if count <= 100 {
				go func(urls []string, ch chan requests.FloatInfo) {

					requests.GetExtraInfo(urls, ch)
					wg.Done()

				}(links[start:start+count], flCh)
				// go requests.GetExtraInfo(links[start:start+count], flCh)
			} else {
				go func(urls []string, ch chan requests.FloatInfo) {

					requests.GetExtraInfo(urls, ch)
					wg.Done()

				}(links[start:start+100], flCh)
				// go requests.GetExtraInfo(links[start:start+100], flCh)
			}
			start += 100
		}
	} else {
		wg.Add(1)
		go func(urls []string, ch chan requests.FloatInfo) {

			requests.GetExtraInfo(urls, ch)
			wg.Done()

		}(links, flCh)
		// go requests.GetExtraInfo(links, flCh)
	}

	go func() {
		wg.Wait()
		fmt.Println("Goroutines done")
		close(flCh)
		fmt.Println("Channel closed")
	}()

	for v := range flCh {
		// fmt.Println(v)
		floatInfoList = append(floatInfoList, v)
	}

	fmt.Println("Elements count: ", len(floatInfoList))

	file, _ := json.MarshalIndent(floatInfoList, "", " ")
	_ = ioutil.WriteFile("floatInfo.json", file, 0644)
	fmt.Println("Full time:", time.Since(startTime))
}
