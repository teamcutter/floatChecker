package main

import (
	"github.com/teamcutter/floatChecker/internal/handlers"
	"github.com/teamcutter/floatChecker/internal/middlewares"

	"github.com/gin-gonic/gin"
)

const thousand = "https://steamcommunity.com/market/listings/730/AK-47%20%7C%20Slate%20%28Minimal%20Wear%29/render/?query=country=EU&language=english&currency=1"
// тут много https://steamcommunity.com/market/listings/730/StatTrak%E2%84%A2%20AK-47%20%7C%20Uncharted%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1
//"https://steamcommunity.com/market/listings/730/StatTrak™%20Desert%20Eagle%20%7C%20Directive%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1"

func main() {
	server := gin.Default()

	server.Use(middlewares.Logger())
	
	server.GET("/api/v1/floats/info/:queryURL", handlers.FloatInfoHandler)
	server.GET("/api/v1/overpriced/:queryURL", handlers.FloatOverpricedHandler)
	server.GET("/api/v1/overpriced/:queryURL/weaponType=:weaponType", handlers.FloatOverpricedByWeaponHandler)


	server.Run(":8080")
}