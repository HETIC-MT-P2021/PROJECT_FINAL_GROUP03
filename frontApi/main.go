package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/router"
)

func main() {
	// start listening on routes
	log.Info("Hey it's the front api")
	r := gin.Default()
	router.Initialize(r)
	r.Run(":8000")
}