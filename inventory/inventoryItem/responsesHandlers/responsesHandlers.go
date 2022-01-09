	package responsesHandlers

	import (
		"fmt"
		"github.com/bwmarrin/snowflake"
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		"github.com/jmoiron/sqlx"
		"joke/config"
		"joke/inventory/inventoryItem"
		"joke/inventory/inventoryItemPo"
		"joke/utils/database"
		"joke/utils/processErr"
		"log"
		"net/http"
		"strconv"
	)
	
// CreateItem localhost:4000/create
/*
{
	"comment": "5555",
	"from_location": "NY",
	"current_location": "NY",
	"to_location": "NY",
	"original_price": 4500,
	"current_price": 860,
	"weight": 34034,
	"url": "6324",
	"name": "sword"
}
*/
func CreateItem(c *gin.Context) {
	log.Println("creating inventory item")

	var item inventoryItem.InventoryItem

	err := c.BindJSON(&item)
	processErr.ProcessInternalErr(c, err, "bindJSON error")

	db := database.OpenDB(c)
	defer db.Close()

	node, err := snowflake.NewNode(1)
	processErr.ProcessInternalErr(c, err, "%v")

	id := node.Generate()

	item.ItemId = strconv.FormatInt(int64(id), 10)

	insert, err1 := db.Prepare(
		"INSERT INTO inventory_item( item_id ,  thumbnail ,  comment , from_location , " +
			" current_location ,  to_location ,original_price ,  current_price ,  weight ,  url, name ) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	_, err2 := insert.Exec(
		item.ItemId,
		item.ThumbNail,
		item.Comment,
		item.FromLocation,
		item.CurrentLocation,
		item.ToLocation,
		item.OriginalPrice,
		item.CurrentPrice,
		item.Weight,
		item.Url,
		item.Name,
	)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "item inserting error",
		})
		panic(err2)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "insert success",
	})
}

// UpdateItem localhost:4000/update?query-id=?
/*
request body:
{
	"comment": "Good job"
}
*/
func UpdateItem(c *gin.Context) {
	log.Println("updating inventory item")

	var item inventoryItem.InventoryItem

	itemId := c.Query("query-id")

	_, err := strconv.Atoi(itemId)
	if itemId == "" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "wrong format of id. Id must be integer",
		})
		panic(err)
		return
	}

	err = c.BindJSON(&item)
	processErr.ProcessInternalErr(c, err, "bindJSON error")

	db := database.OpenDB(c)
	defer db.Close()

	update, err1 := db.Prepare(
		"UPDATE   inventory_item   SET " +
			"  item_id   = ?,   thumbnail   = ?,   comment   = ?, " +
			"  from_location   = ?,   current_location   = ?,   to_location   = ?, " +
			"  original_price   = ?,   current_price   = ?,   weight   = ?,   url   = ?, " +
			"  name   = ? WHERE   id   = ?;",
	)

	_, err2 := update.Exec(
		item.ItemId,
		item.ThumbNail,
		item.Comment,
		item.FromLocation,
		item.CurrentLocation,
		item.ToLocation,
		item.OriginalPrice,
		item.CurrentPrice,
		item.Weight,
		item.Url,
		item.Name,
		itemId,
	)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "item update error",
		})
		panic(err2)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}

// DeleteItem localhost:4000/delete?query-id=?
func DeleteItem(c *gin.Context) {
	log.Println("deleting inventory item")

	itemId := c.Query("query-id")

	db := database.OpenDB(c)
	defer db.Close()

	deletion, err1 := db.Prepare(
		"DELETE FROM inventory_item where id = ?",
	)

	_, err2 := deletion.Exec(itemId)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "item deleting error",
		})
		panic(err2)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})
}

// GetItemList localhost:4000/get-list?start=?&size=?
func GetItemList(c *gin.Context) {
	log.Println("get inventory items based on pagination")

	s := c.Query("start")
	p := c.Query("size")

	start, err1 := strconv.Atoi(s)
	pageSize, err2 := strconv.Atoi(p)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "wrong format of id. Id must be integer",
		})
		panic(err1)
		return
	}

	log.Println(fmt.Sprintf("pageSize: %d, start: %d", pageSize, (start - 1) * pageSize))

	db, err := sqlx.Connect("mysql", config.DBConnection)
	processErr.ProcessInternalErr(c, err, "mysql opening error")
	defer db.Close()

	var result []inventoryItemPo.InventoryItemPo

	err = db.Select(&result,
		"select item_id ,  thumbnail ,  comment , from_location , " +
		" current_location ,  to_location ,original_price ,  current_price ,  weight ,  " +
		"url, name, id from inventory_item limit ? offset ?", pageSize, (start - 1) * pageSize)
	processErr.ProcessInternalErr(c, err, "inventory list get error")

	c.JSON(http.StatusOK, gin.H {
		"list": result,
	})
}
