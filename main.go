package main

import (
	"log"
	"os"
	"restaurantserver/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
    port := os.Getenv("PORT")
	router := gin.New()

	router.Use(gin.Logger())

    router.Use(cors.Default())

	//Api endpoints
	router.GET("/",routes.Welcome)
    router.POST("/order/create",routes.AddOrder)

	router.GET("/orders",routes.GetOrders)

	router.PUT("/order/update/:id",routes.UpdateOrder)

	router.DELETE("/order/delete/:id",routes.DeleteOrder)

	//running the server
	router.Run(":"+port)
	
}