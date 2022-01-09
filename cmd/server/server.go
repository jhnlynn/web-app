package main

import (
	"web-app/routers"
)

func main() {

	router := routers.Routers()

	_ = router.Run(":4000")
}

