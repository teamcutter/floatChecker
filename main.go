package main

import (
	"floatChecker/internal/requests"

	"github.com/gofiber/fiber/v2"
)

const thousand = "https://steamcommunity.com/market/listings/730/AK-47%20%7C%20Slate%20%28Minimal%20Wear%29/render/?query=country=EU&language=english&currency=1"
// тут много https://steamcommunity.com/market/listings/730/StatTrak%E2%84%A2%20AK-47%20%7C%20Uncharted%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1

func main() {
	app := fiber.New()
	url := "https://steamcommunity.com/market/listings/730/StatTrak™%20Desert%20Eagle%20%7C%20Directive%20%28Field-Tested%29/render/?query=country=EU&language=english&currency=1"
	
	app.Get("/", func(ctx *fiber.Ctx) error {
		links := requests.SearchCurrentItem(url)
		floatInfoList := requests.InfoCurrentItem(links)
		return ctx.JSON(floatInfoList)
	})

	app.Listen(":3000")
}
