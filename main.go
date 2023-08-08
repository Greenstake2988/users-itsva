package main

import (
	"users-itsva/database"
	"users-itsva/middlewares"
	"users-itsva/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create db object
	d := &database.DB{}
	// Connect db object to Database
	d.ConnectDB()

	// Create router 
	router := gin.Default()

	// Amake router use the function that Add Database to the context of gin
	router.Use(middlewares.AddDBToContext(d.DB))

	// Setuo the user routes in the router
	routes.SetupUserRoutes(router)

	router.Run()
}
