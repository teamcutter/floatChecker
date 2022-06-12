package requests

import (
	"fmt"
	"math"
	"sync"
	"io"
	"net/http"
	"time"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

const floatUrl string = "https://api.csgofloat.com/?url="

func SearchCurrentItem(url string) []string {

	startTime := time.Now()

	url = url + "&count=100"
	myClient := &http.Client{}
	res, _ := myClient.Get(url)

	body, _ := io.ReadAll(res.Body)

	defer res.Body.Close()

	start := 0

	itemsCount, _ := strconv.Atoi(gjson.Get(string(body), "total_count").String())

	pageCount := int(math.Ceil(float64(itemsCount) / 100))

	var skinUrl string
	var links []string

	for i := 0; i < pageCount; i++ {
		count := itemsCount - start
		if count <= 100 {
			skinUrl = url + "&start=" + strconv.Itoa(start) + "&count=" + strconv.Itoa(count)
		} else {
			skinUrl = url + "&start=" + strconv.Itoa(start) + "&count=100"
		}

		myClient := &http.Client{}
		res, _ := myClient.Get(skinUrl)

		body, _ := io.ReadAll(res.Body)

		defer res.Body.Close()

		data := gjson.Get(string(body), "listinginfo")                       // get raw JSON
		dataString := []byte(data.String())                                  // present it as byteArray
		newDataString := "[" + string(dataString[1:len(dataString)-1]) + "]" // replace first { and last } -> []

		listingIdArray := gjson.Get(newDataString, "#.listingid").Array()
		assetIdArray := gjson.Get(newDataString, "#.asset.id").Array()
		rawLinksArray := gjson.Get(newDataString, "#.asset.market_actions.0.link").Array()
		/* price := gjson.Get(newDataString, "#.converted_price").Array()

		for _, value := range price {
			fmt.Println(value.Float() / 100)
		}
		fmt.Println(price) */

		for i := 0; i < len(listingIdArray); i++ {

			link := strings.Replace(rawLinksArray[i].String(), "%listingid%", listingIdArray[i].String(), 1)
			link = strings.Replace(link, "%assetid%", assetIdArray[i].String(), 1)
			links = append(links, link)
		}

		start += 100
	}
	end := time.Now()
	fmt.Println("End: ", end.Sub(startTime))

	return links // возвращаем только линки, больше нам ничего не нужно по идее
}

func GetExtraInfo(urls []string, ch chan FloatInfo) {
	startTime := time.Now()
	myClient := &http.Client{}
	fmt.Println("Started goroutine")

	for i := 0; i < len(urls); i++ {
		res, _ := myClient.Get(floatUrl + urls[i])

		body, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		stickersJSON := gjson.Get(string(body), "iteminfo.stickers.#.name").Array()
		var stickers []string

		for _, sticker := range stickersJSON {
			stickers = append(stickers, sticker.String())
		}

		ch <- FloatInfo{
			FullItemName: gjson.Get(string(body), "iteminfo.full_item_name").String(),
			FloatValue:   gjson.Get(string(body), "iteminfo.floatvalue").Float(),
			Stickers:     stickers,
		}

	}
	end := time.Now()
	fmt.Println("End: ", end.Sub(startTime))
}

func InfoCurrentItem(links []string) []FloatInfo {
	
	var wg sync.WaitGroup

	flCh := make(chan FloatInfo)
	var floatInfoList []FloatInfo

	start := 0

	countOfGoRoutines := int(math.Ceil(float64(len(links)) / 100))
	fmt.Println("Count of goroutines: ", countOfGoRoutines)

	if len(links) > 100 {
		for i := 0; i < countOfGoRoutines; i++ {
			wg.Add(1)
			count := len(links) - start
			if count <= 100 {
				go func(urls []string, ch chan FloatInfo) {

					GetExtraInfo(urls, ch)
					wg.Done()
					fmt.Println("Done goroutine")

				}(links[start:start+count], flCh)
			} else {
				go func(urls []string, ch chan FloatInfo) {

					GetExtraInfo(urls, ch)
					wg.Done()
					fmt.Println("Done goroutine")

				}(links[start:start+100], flCh)
			}
			start += 100
		}
	} else {
		wg.Add(1)
		go func(urls []string, ch chan FloatInfo) {

			GetExtraInfo(urls, ch)
			wg.Done()

		}(links, flCh)
	}

	// we need to wait for all goroutines to finish at the same time while they are working 
	go func() {
		wg.Wait()
		fmt.Println("Goroutines done")
		close(flCh)
		fmt.Println("Channel closed")
	}()

	for v := range flCh {
		floatInfoList = append(floatInfoList, v)
	}

	return floatInfoList
}