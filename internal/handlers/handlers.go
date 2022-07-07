package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/teamcutter/floatChecker/internal/search"
	"github.com/teamcutter/floatChecker/internal/entities"

	"github.com/gin-gonic/gin"
)

const baseURLOverpriced = "https://inventories.cs.money/5.0/load_bots_inventory/730?"

func FloatInfoHandler(c *gin.Context) {

	log.Println("queryURL: ", c.Param("queryURL"))
	url := fmt.Sprintf("https://steamcommunity.com/market/listings/730/%s/render/?query=country=EU&language=english&currency=1", c.Param("queryURL"))
	log.Println("Requesting: ", url)
	links := search.SearchCurrentItem(url)
	floatInfoList := search.InfoCurrentItem(links)
	sort.SliceStable(floatInfoList, func(i, j int) bool {
		return floatInfoList[i].FloatValue < floatInfoList[j].FloatValue
	})
	c.IndentedJSON(http.StatusOK, floatInfoList)
}

func FloatOverpricedHandler(c *gin.Context) {

	var items []entities.OverpricedItem
	if c.Param("save") == "true" {

		items = search.OverpricedInfo(baseURLOverpriced + c.Param("queryURL"), c.Param("save"), "all", nil)
	} else if c.Param("weaponType") != "all" {

		items = search.OverpricedInfo(baseURLOverpriced + c.Param("queryURL"), "false", c.Param("weaponType"), 
		func(item entities.OverpricedItem, wt string) bool {
			return strings.Contains(item.FullName, strings.ToUpper(wt))
		})
	} else if c.Param("save") == "true" && c.Param("weaponType") != "all" {

		items = search.OverpricedInfo(baseURLOverpriced + c.Param("queryURL"), c.Param("save"), c.Param("weaponType"), 
		func(item entities.OverpricedItem, wt string) bool {
			return strings.Contains(item.FullName, strings.ToUpper(wt))
		})
	} else {

		items = search.OverpricedInfo(baseURLOverpriced + c.Param("queryURL"), "false", "all", nil)
	}
	
	c.IndentedJSON(http.StatusOK, items)
}