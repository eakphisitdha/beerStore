package main

import (
	"beerstore/database"
	"beerstore/handler"
	"beerstore/repository"
	"beerstore/service"
	"beerstore/transaction"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Mariadb()
	defer db.Close()
	conn := database.MongoDB()
	defer conn.Client().Disconnect(context.Background())

	r := repository.NewRepository(db)
	t := transaction.NewTransaction(conn)
	s := service.NewService(r, t)
	h := handler.NewHandler(s)

	router := gin.Default()

	router.GET("/beer", h.Get)
	router.POST("/beer", h.Add)
	router.PUT("/beer/:id", h.Update)
	router.DELETE("/beer/:id", h.Delete)

	if err := router.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
