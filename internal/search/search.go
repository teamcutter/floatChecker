package search

import (
	"github.com/teamcutter/floatChecker/internal/entities"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/tidwall/gjson"
)

const floatUrl string = "https://api.csgofloat.com/?url="

func SearchCurrentItem(url string) []string {

	url = url + "&count=100"
	myClient := &http.Client{}
	res, err := myClient.Get(url)
	if err != nil {
		log.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	start := 0

	itemsCount, err := strconv.Atoi(gjson.Get(string(body), "total_count").String())
	if err != nil {
		log.Println(err)
	}
	log.Printf("Expected items count: %d \n", itemsCount)

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

		res, err := myClient.Get(skinUrl)
		if err != nil {
			log.Println(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}

		data := gjson.Get(string(body), "listinginfo")                       // get raw JSON
		dataString := []byte(data.String())                                  // present it as byteArray
		newDataString := "[" + string(dataString[1:len(dataString)-1]) + "]" // replace first { and last } -> []

		listingIdArray := gjson.Get(newDataString, "#.listingid").Array()
		assetIdArray := gjson.Get(newDataString, "#.asset.id").Array()
		rawLinksArray := gjson.Get(newDataString, "#.asset.market_actions.0.link").Array()
		/* price := gjson.Get(newDataString, "#.converted_price").Array()

		for _, value := range price {
			log.Println(value.Float() / 100)
		}
		log.Println(price) */

		for i := 0; i < len(listingIdArray); i++ {

			link := strings.Replace(rawLinksArray[i].String(), "%listingid%", listingIdArray[i].String(), 1)
			link = strings.Replace(link, "%assetid%", assetIdArray[i].String(), 1)
			links = append(links, link)
		}

		start += 100
	}

	log.Printf("Real items count: %d \n", len(links))

	return links // возвращаем только линки, больше нам ничего не нужно по идее
}

func GetExtraInfo(urls []string, ch chan entities.FloatInfo) {

	myClient := &http.Client{}

	for i := 0; i < len(urls); i++ {
		res, err := myClient.Get(floatUrl + urls[i]); if err != nil {
			log.Println(err)
		}

		body, err := io.ReadAll(res.Body); if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()

		stickersJSON := gjson.Get(string(body), "iteminfo.stickers.#.name").Array()
		var stickers []string

		for _, sticker := range stickersJSON {
			stickers = append(stickers, sticker.String())
		}

		ch <- entities.FloatInfo{
			FullItemName: gjson.Get(string(body), "iteminfo.full_item_name").String(),
			FloatValue:   gjson.Get(string(body), "iteminfo.floatvalue").Float(),
			Stickers:     stickers,
		}

	}
}

func InfoCurrentItem(links []string) []entities.FloatInfo {

	var wg sync.WaitGroup

	flCh := make(chan entities.FloatInfo)
	var floatInfoList []entities.FloatInfo

	start := 0

	countOfGoRoutines := int(math.Ceil(float64(len(links)) / 50))
	log.Printf("%d goroutines working \n", countOfGoRoutines)
	log.Println("Collecting...")

	if len(links) > 50 {
		for i := 0; i < countOfGoRoutines; i++ {
			wg.Add(1)
			count := len(links) - start
			if count <= 50 {
				go func(urls []string, ch chan entities.FloatInfo) {

					GetExtraInfo(urls, ch)
					wg.Done()

				}(links[start:start+count], flCh)
			} else {
				go func(urls []string, ch chan entities.FloatInfo) {

					GetExtraInfo(urls, ch)
					wg.Done()

				}(links[start:start+50], flCh)
			}
			start += 50
		}
	} else {
		wg.Add(1)
		go func(urls []string, ch chan entities.FloatInfo) {

			GetExtraInfo(urls, ch)
			wg.Done()

		}(links, flCh)
	}

	// we need to wait for all goroutines to finish at the same time while they are working
	go func() {
		wg.Wait()
		close(flCh)
		log.Println("Data collected!")
	}()

	for v := range flCh {
		floatInfoList = append(floatInfoList, v)
	}

	return floatInfoList
}
