package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/outstream-agg/api/routes"
	util "github.com/micro/outstream-agg/utils"
	"net/http"
	"strings"
	"time"
)

func main() {
	router := gin.Default()

	// Bind routes
	routes.ApplicationV1Router(router)
	startServer(router)
}

func startServer(router http.Handler) {
	s := &http.Server{
		Addr:         util.GetConfig("ServerPort"),
		Handler:      router,
		ReadTimeout:  18000 * time.Second,
		WriteTimeout: 18000 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		_ = fmt.Errorf("Error while starting the server %s", strings.ToLower(err.Error()))
		panic(err)
	}
}
