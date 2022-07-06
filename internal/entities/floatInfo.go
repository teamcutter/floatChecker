package entities


type FloatInfo struct {
	FullItemName string    `json:"full_item_name"`
	FloatValue  float64 `json:"float_value"`
	Stickers     []string  `json:"stickers"`
}

type OverpricedItem struct {
	AppID         int           `json:"appId"`
	AssetID       int64         `json:"assetId"`
	Float         string        `json:"float"`
	HasHighDemand bool          `json:"hasHighDemand"`
	ID            int64         `json:"id"`
	Img           string        `json:"img"`
	NameID        int           `json:"nameId"`
	Overprice     float64       `json:"overprice"`
	Price         float64       `json:"price"`
	Quality       string        `json:"quality"`
	Rarity        string        `json:"rarity"`
	SteamID       string        `json:"steamId"`
	SteamImg      string        `json:"steamImg"`
	Stickers      []interface{} `json:"stickers"`
	Type          int           `json:"type"`
	UserID        interface{}   `json:"userId"`
	Pattern       int           `json:"pattern"`
	Rank          interface{}   `json:"rank"`
	Collection    string        `json:"collection"`
	Overpay       struct {
		Float    float64 `json:"float"`
		Stickers float64     `json:"stickers"`
	} `json:"overpay"`
	Inspect  string `json:"inspect"`
	FullName string `json:"fullName"`
}