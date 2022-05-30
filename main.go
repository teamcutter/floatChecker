package main

import (
	//"encoding/json"
	"floatChecker/requests"
	"fmt"
	_"io"
	_ "math"
	_ "strings"
)

func main() {
	url := "https://steamcommunity.com/market/listings/730/StatTrak%E2%84%A2%20AK-47%20%7C%20Uncharted%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1"
	itemsList := requests.SearchCurrentItem(url)
	fmt.Println(len(itemsList))
}