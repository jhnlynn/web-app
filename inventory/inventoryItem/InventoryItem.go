package inventoryItem

// InventoryItem
// ItemId: uuid
// ThumbNail: thumbnail
// Comment: comment on the item
// FromLocation: location the item delivered from
// (if is still in inventory, this field should be null)
// CurrentLocation: location the item currently is
// ToLocation: destination
// OriginalPrice: item's original price
// CurrentPrice: item's current price
// Weight: item's weight
// Url: item's URL
type InventoryItem struct {
	ItemId string `db:"item_id" json:"item_id"`
	ThumbNail string `db:"thumbnail" json:"thumb_nail"`
	Comment string `db:"comment" json:"comment"`
	FromLocation string `db:"from_location" json:"from_location"`
	CurrentLocation string `db:"current_location" json:"current_location"`
	ToLocation string `db:"to_location" json:"to_location"`
	OriginalPrice int32 `db:"original_price" json:"original_price"`
	CurrentPrice int32 `db:"current_price" json:"current_price"`
	Weight int32 `db:"weight" json:"weight"`
	Url string `db:"url" json:"url"`
	Name string `db:"name" json:"name"`
}
