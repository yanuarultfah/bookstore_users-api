package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func startApplication() {
	mapUrls()
	router.Run(addr, ":8080")
}
