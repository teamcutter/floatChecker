package handlers

import (
	"github.com/teamcutter/floatChecker/internal/search"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)


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

}