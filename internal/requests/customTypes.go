package requests


type FloatInfo struct {
	FullItemName string    `json:"full_item_name"`
	FloatValue  float64 `json:"floats"`
	Stickers     []string  `json:"stickers"`
}
