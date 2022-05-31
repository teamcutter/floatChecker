package requests

import (
	_ "encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "strconv"
	_ "strings"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

const floatUrl string = "https://api.csgofloat.com/?url="

func GetExtraInfo(urls []string, ch chan FloatInfo, wg *sync.WaitGroup) /* []FloatInfo */ {
	startTime := time.Now()
	myClient := &http.Client{}
	fmt.Println("Started goroutine")
	defer wg.Done()

	// var floatInfoList []FloatInfo
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
			FloatValue:   gjson.Get(string(body), "iteminfo.floats").Float(),
			Stickers:     stickers,
		}
		/* if ch != nil {
			ch <- FloatInfo{
				FullItemName: gjson.Get(string(body), "iteminfo.full_item_name").String(),
				FloatValue:   gjson.Get(string(body), "iteminfo.floats").Float(),
				Stickers:     stickers,
			}
		} else {
			floatInfoList = append(floatInfoList, FloatInfo{
				FullItemName: gjson.Get(string(body), "iteminfo.full_item_name").String(),
				FloatValue:   gjson.Get(string(body), "iteminfo.floats").Float(),
				Stickers:     stickers,
			})
		} */

		/* floatInfoList = append(floatInfoList, FloatInfo{
			FullItemName: gjson.Get(string(body), "iteminfo.full_item_name").String(),
			FloatValue:   gjson.Get(string(body), "iteminfo.floatvalue").Float(),
			Stickers:     stickers,
		}) */

	}
	// ch <- floatInfoList
	end := time.Now()
	fmt.Println("End: ", end.Sub(startTime))

	// return floatInfoList
}
