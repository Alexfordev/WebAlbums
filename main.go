package main

import (
	"github.com/alexfordev/WebAlbums/log"
	"github.com/alexfordev/WebAlbums/routers"
)

func main() {
	router := routers.InitRouter()
	log.Log.Debug("main in")
	_ = router.Run(":8080")
}
