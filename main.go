package main

import (
	"encoding/json"
	"floatChecker/requests"
	"fmt"
	_ "io"
	"io/ioutil"
	_ "math"
	_ "strings"
	"time"
	"sync"
	"math"
)

const thousand = "https://steamcommunity.com/market/listings/730/AK-47%20%7C%20Slate%20%28Minimal%20Wear%29/render/?query=country=EU&language=english&currency=1"

func main() {
	startTime := time.Now()
	// тут много https://steamcommunity.com/market/listings/730/StatTrak%E2%84%A2%20AK-47%20%7C%20Uncharted%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1
	url := "https://steamcommunity.com/market/listings/730/AK-47%20%7C%20Elite%20Build%20%28Battle-Scarred%29/render/?query=country=EU&language=english&currency=1"
	links := requests.SearchCurrentItem(url)
	fmt.Println("Elements: ", len(links))

	var wg sync.WaitGroup

	flCh := make(chan requests.FloatInfo)

	var floatInfoList []requests.FloatInfo

	start := 0

	if len(links) > 100 {
		// 130 -> 100 + 30 
		countOfGoRoutines := int(math.Ceil(float64(len(links)) / 100))
		fmt.Println("Count of goroutines: ", countOfGoRoutines)
		for i := 0; i < countOfGoRoutines; i++ {

			count := len(links) - start
			wg.Add(1)
			if count <= 100 {
				go requests.GetExtraInfo(links[start:start+count], flCh, &wg)	
				
			} else {
				go requests.GetExtraInfo(links[start:start+100], flCh, &wg)
			}
			start += 100
		} 
	} else {
		wg.Add(1)
		go requests.GetExtraInfo(links, flCh, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("Goroutines done")
		close(flCh)
	}()

	for v := range flCh {
		floatInfoList = append(floatInfoList, v)
	}

	file, _ := json.MarshalIndent(floatInfoList, "", " ")
	_ = ioutil.WriteFile("floatInfo.json", file, 0644)
	fmt.Println("Full time:", time.Since(startTime))
}
