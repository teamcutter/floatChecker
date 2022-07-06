package search

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/teamcutter/floatChecker/internal/entities"
)

// url https://inventories.cs.money/5.0/load_bots_inventory/730?buyBonus=35&hasRareFloat=true&isStore=true&limit=60&maxPrice=10000&minPrice=1&offset=5000&sort=botFirst&type=5&type=6&type=3&type=4&type=7&type=8&withStack=true

func OverpricedInfo(url, save, weaponType string, filter func(entities.OverpricedItem, string) bool) []entities.OverpricedItem {

	myClient := &http.Client{}
	var itemsJSON map[string][]entities.OverpricedItem
	var items []entities.OverpricedItem

	offsetCount := 0
	for {
		log.Printf("Offset: %d\n", offsetCount)
		res, err := myClient.Get(url + fmt.Sprintf("&offset=%d", offsetCount)); if err != nil {
			log.Println(err)
		} 
		
		body, err := io.ReadAll(res.Body); if err != nil {
			log.Println(err)
		}
		
		err = json.Unmarshal(body, &itemsJSON)
		if err != nil {
			log.Println(err)
			defer res.Body.Close()
			break
		}
		if itemsList, ok := itemsJSON["items"]; ok {
			items = append(items, itemsList...)
		} else {
			defer res.Body.Close()
			break
		}

		offsetCount += 50
	}
	if weaponType != "all" {
		var filteredItems []entities.OverpricedItem
		for _, item := range items {
			if filter(item, weaponType) {
				filteredItems = append(filteredItems, item)
			}
		}
		if save == "true" {
			filteredItemsJSON, err := json.Marshal(filteredItems); if err != nil {
				log.Println(err)
			}
			err = ioutil.WriteFile(weaponType + ".json", filteredItemsJSON, 0644); if err != nil {
				log.Println(err)
			}
		}
		return filteredItems
	}
	if save == "true" {
		itemsJSON, err := json.Marshal(items); if err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile("db.json", itemsJSON, 0644); if err != nil {
			log.Println(err)
		}
	}
	return items
}