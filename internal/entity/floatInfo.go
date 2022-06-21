package entity


type FloatInfo struct {
	FullItemName string    `json:"full_item_name"`
	FloatValue  float64 `json:"float_value"`
	Stickers     []string  `json:"stickers"`
}
