package requests

/* type ItemInfo struct {
	ListingId string `json:listingid`
	AssetId string `json:assetid`
	Link string `json:link`
} */

type FloatInfo struct {
	FullItemName string    `json:full_item_name`
	FloatValue  float64 `json:floats`
	Stickers     []string  `json:stickers`
}
