package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/ratnesh-maurya/api.ratn.tech/config"
	"github.com/ratnesh-maurya/api.ratn.tech/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectMongoDB()
	corsConfig := cors.Config{
		Origins:         "*",
		RequestHeaders:  "Origin, Authorization, Content-Type,App-User, Org_id",
		Methods:         "GET, POST, PUT,DELETE",
		Credentials:     false,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	gin.SetMode(gin.ReleaseMode)
router := gin.New() 
router.Use(cors.Middleware(corsConfig))
router.Use(gin.Logger(), gin.Recovery()) 

	routes.RatnTechRoutes(router)

log.Print("Server listening on http://localhost:8000/")
	if err := http.ListenAndServe("0.0.0.0:8000", router); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
	
}
