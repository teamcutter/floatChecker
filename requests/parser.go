package requests

import (
	"fmt"
	"io"
	_ "io"
	_ "math"
	"net/http"
	_ "strings"
	_"time"

	"github.com/tidwall/gjson"
)

func SearchCurrentItem(url string) {

	myClient := &http.Client{}
	res, err := myClient.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	data := gjson.Get(string(body), "listinginfo")
	gjson.ForEachLine(data.String(), func(line gjson.Result) bool {
		fmt.Println(line.String())
		// fmt.Println(gjson.Get(string(line.Raw), "listingid"))
		return true
	})
	// fmt.Println(data)
	defer res.Body.Close()
}
