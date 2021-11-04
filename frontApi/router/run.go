package router

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine) {
	configureCORS(r)

	go func() {
		if err := r.Run(fmt.Sprintf(":%s", env.GetVariable("PORT"))); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// ----------------- CLOSE APP -----------------
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	consume(quit)
	log.Info("Shutting down server...")
}

func consume(ch <-chan os.Signal) os.Signal {
	return <-ch
}

func configureCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"},
		ExposeHeaders:    []string{"Authorization"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "User-Agent", "Referrer", "Host"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
