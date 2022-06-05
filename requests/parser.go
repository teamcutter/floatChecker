package requests

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

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
