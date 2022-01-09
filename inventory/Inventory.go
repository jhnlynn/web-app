package inventory

// Inventory
// InventoryId: uuid
// Location: inventory location
// Comment: inventory comment
// InventoryAdmin: administrator of the inventory
// Capacity: capacity of inventory
type Inventory struct {
	InventoryId string `json:"inventory_id"`
	Location string `json:"location"`
	Comment string `json:"comment"`
	InventoryAdmin string `json:"inventory_admin"`
	Capacity int32 `json:"capacity"`
	Name string `json:"name"`
}
