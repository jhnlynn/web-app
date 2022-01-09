package main

import (
	"joke/routers"
)

func main() {

	router := routers.Routers()

	_ = router.Run(":4000")
}

