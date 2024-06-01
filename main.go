package main

import (

	"github.com/gin-gonic/gin"
	"postgres-metadata/db"
	"postgres-metadata/handlers"
	"postgres-metadata/repository"
	"postgres-metadata/services"
)

func main() {

	dbConn := db.ConnectDatabase("postgres://postgres:123@localhost/postgres_metadata?sslmode=disable")

	router := gin.Default()
	
	repo := repository.NewPostgresRepository(dbConn)
	service := services.NewMetadataService(repo)
	handler := handlers.NewMetadataHandler(service)

	router.GET("/tables/:schema", handler.GetTables)
	router.GET("/columns/:schema/:table", handler.GetColumns)
	router.GET("/indexes/:schema", handler.GetIndexes)
	router.GET("/constraints/:schema/:table", handler.GetConstraints)
	router.GET("/schemas", handler.GetSchemas)
	router.GET("/views/:schema", handler.GetViews)
	router.GET("/sequences/:schema", handler.GetSequences)
	router.GET("/functions/:schema", handler.GetFunctions)
	router.GET("/triggers/:schema", handler.GetTriggers)
	router.GET("/roles", handler.GetRoles)

	router.Run("localhost:8080")
}
