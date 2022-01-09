package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"joke/inventory/inventoryItem/responsesHandlers"
	"joke/utils/imageUpload"

	//"go-play/middleware"
)

func Routers() *gin.Engine {

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		cors.Default()
		c.Next()
		gin.Logger()
		gin.Recovery()
		//middleware.CORSMiddleware()
	})

	// upload image localhost:4000/upload_test?item-id=?
	router.POST("/upload_test", imageUpload.UploadImageAndThumb)

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
	router.POST("/create", responsesHandlers.CreateItem)

	// UpdateItem localhost:4000/update?query-id=?
	/*
	   request body:
	   {
	   	"comment": "Good job"
	   }
	*/
	router.POST("/update", responsesHandlers.UpdateItem)

	// DeleteItem localhost:4000/delete?query-id=?
	router.GET("/delete", responsesHandlers.DeleteItem)

	// GetItemList localhost:4000/get-list?start=?&size=? [start > 0, size > 0]
	router.GET("/get-list", responsesHandlers.GetItemList)

	return router
}
