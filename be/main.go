package main

import (
	"context"
	"fmt"
	"main/src/api"
	"main/src/database"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	println("starting")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic %v \n call stack\n %s\n\nexit", err, debug.Stack())
		}

		if database.DB != nil {
			db, err := database.DB.DB()
			if err == nil {
				db.Close()
			}
		}
	}()
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	// some router
	// token
	router.POST("token", api.PostToken)
	router.DELETE("token", api.DeleteToken)
	router.PUT("token", api.PutToken)
	// user
	router.GET("user", api.GetUser)
	router.POST("user", api.PostUser)
	router.PATCH("user", api.PatchUser)
	router.DELETE("user", api.DeleteUser)

	server := &http.Server{
		Addr:    ":3623",
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	tool := database.NewDebug()
	println("server started")
	quit := make(chan os.Signal, 5)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tool.Stop(ctx)
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
	println("server exit")
}
