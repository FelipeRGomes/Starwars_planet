package controller

import (
	"statwars_planets/entity"
	"statwars_planets/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type PlanetController interface {
	Save(ctx *gin.Context, s *mgo.Session) error
	Delete(ctx *gin.Context, s *mgo.Session) error
	FindAll(s *mgo.Session) []entity.Planet
	Find(ctx *gin.Context, s *mgo.Session) []entity.Planet
}

type controller struct {
	service service.PlanetService
}

func New(service service.PlanetService) PlanetController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(s *mgo.Session) []entity.Planet {
	return c.service.FindAll(s)
}

func (c *controller) Find(ctx *gin.Context, s *mgo.Session) []entity.Planet {
	var ps entity.PlanetSearch
	err := ctx.ShouldBindJSON(&ps)
	if err != nil {
		panic(err)
	}
	return c.service.Find(ps, s)
}

func (c *controller) Save(ctx *gin.Context, s *mgo.Session) error {
	var planet entity.Planet
	err := ctx.ShouldBindJSON(&planet)
	if err != nil {
		return err
	}

	planet.FilmCount = c.service.MovieCount(ctx, planet)

	c.service.Save(planet, s)
	return nil
}

func (c *controller) Delete(ctx *gin.Context, s *mgo.Session) error {
	var ps entity.PlanetSearch
	err := ctx.ShouldBindJSON(&ps)
	if err != nil {
		return err
	}

	c.service.Delete(ps, s)
	return nil
}
