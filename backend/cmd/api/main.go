package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"

	"github.com/fiorellizz/gopayflow/internal/application"
	"github.com/fiorellizz/gopayflow/internal/infrastructure/database"
	httpInterface "github.com/fiorellizz/gopayflow/internal/interfaces/http"
)

func main() {

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	orderRepo := database.NewPostgresOrderRepository(db)

	createOrderUseCase := application.NewCreateOrderUseCase(orderRepo)
	getOrderByIDUseCase := application.NewGetOrderByIDUseCase(orderRepo)
	listOrdersUseCase := application.NewListOrdersUseCase(orderRepo)

	orderHandler := httpInterface.NewOrderHandler(
		createOrderUseCase,
		getOrderByIDUseCase,
		listOrdersUseCase,
	)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running",
		})
	})

	router.POST("/orders", orderHandler.CreateOrder)
	router.GET("/orders", orderHandler.ListOrders)
	router.GET("/orders/:id", orderHandler.GetOrderByID)

	log.Println("Starting server on port 8080")
	router.Run(":8080")
}