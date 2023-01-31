package main

import (
	"github.com/gin-gonic/gin"
	"wink/router"
	"wink/service"
)

func main() {
	r := gin.Default()
	db := service.InitDB()
	print(db)
	r = router.InitRouter(r)
	panic(r.Run(":9090"))
}
