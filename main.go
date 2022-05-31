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
)

func main() {
	startTime := time.Now()
	url := "https://steamcommunity.com/market/listings/730/StatTrak%E2%84%A2%20AK-47%20%7C%20Uncharted%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1"
	links := requests.SearchCurrentItem(url)
	fmt.Println("Elements: ", len(links))

	floatInfo := requests.GetExtraInfo(links)
	fmt.Println(floatInfo)
	
	file, _ := json.MarshalIndent(floatInfo, "", " ")
	_ = ioutil.WriteFile("floatInfo.json", file, 0644)
	fmt.Println("Full time:", time.Since(startTime))
}