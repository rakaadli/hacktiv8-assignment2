package main

import (
	"hacktiv8-assignment2/controllers"
	"hacktiv8-assignment2/database"
	"hacktiv8-assignment2/repositories"
	"hacktiv8-assignment2/services"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	APP_PORT = ":3000"
)

func main() {

	db := database.ConnectDB()
	router := gin.Default()

	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(orderRepo)

	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	orderController := controllers.NewOrderController(orderService, itemService)

	router.POST("/orders", orderController.CreateNewOrder)
	router.GET("/orders", orderController.GetAllOrdersWithItems)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	log.Println("Server running ar port : ", APP_PORT)
	router.Run(APP_PORT)
}
