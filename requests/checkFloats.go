package requests

import (
	_ "encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "strconv"
	_ "strings"
	"time"

	"github.com/tidwall/gjson"
)

const floatUrl string = "https://api.csgofloat.com/?url="

func GetExtraInfo(urls []string) []FloatInfo {

	startTime := time.Now()
	myClient := &http.Client{}

	var floatInfoList []FloatInfo
	// test := "https://api.csgofloat.com/?url="
	for i := 0; i < len(urls); i++ {
		res, _ := myClient.Get(floatUrl + urls[i])

		body, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		stickersJSON := gjson.Get(string(body), "iteminfo.stickers.#.name").Array()
		var stickers []string

		for _, sticker := range stickersJSON {
			stickers = append(stickers, sticker.String())
		}

		floatInfoList = append(floatInfoList, FloatInfo{
			FullItemName: gjson.Get(string(body), "iteminfo.full_item_name").String(),
			FloatValue:   gjson.Get(string(body), "iteminfo.floatvalue").Float(),
			Stickers:     stickers,
		})

	}

	end := time.Now()
	fmt.Println("End: ", end.Sub(startTime))

	return floatInfoList
}
