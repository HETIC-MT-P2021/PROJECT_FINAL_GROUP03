package main

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// start the router
	r := gin.Default()
	router.Initialize(r)
	router.Run(r)
}