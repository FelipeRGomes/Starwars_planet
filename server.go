package main

import (
	"net/http"
	"statwars_planets/controller"
	"statwars_planets/service"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

var (
	planetService    service.PlanetService       = service.New()
	planetController controller.PlanetController = controller.New(planetService)
)

func main() {
	server := gin.Default()

	session, err := mgo.Dial("db")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// MongoDB open connection mode for multiple reads (?)
	session.SetMode(mgo.Monotonic, true)

	server.GET("/planets/all", func(ctx *gin.Context) {
		ctx.JSON(200, planetController.FindAll(session))
	})
	server.GET("/planets", func(ctx *gin.Context) {
		ctx.JSON(200, planetController.Find(ctx, session))
	})

	server.POST("/planets", func(ctx *gin.Context) {
		err := planetController.Save(ctx, session)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Planet Registered."})
		}

	})
	server.DELETE("/planets", func(ctx *gin.Context) {
		err := planetController.Delete(ctx, session)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Planet Deleted."})
		}

	})
	server.Run(":3000")

}
