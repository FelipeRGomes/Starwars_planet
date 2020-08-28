package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"statwars_planets/entity"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const banco = "starwars"

type PlanetService interface {
	Save(entity.Planet, *mgo.Session) error
	Delete(entity.PlanetSearch, *mgo.Session) error
	FindAll(*mgo.Session) []entity.Planet
	Find(entity.PlanetSearch, *mgo.Session) []entity.Planet
	MovieCount(*gin.Context, entity.Planet) int
}

type planetService struct {
	planets []entity.Planet
}

func New() PlanetService {
	return &planetService{
		planets: []entity.Planet{},
	}
}

func (service *planetService) Save(planet entity.Planet, s *mgo.Session) error {
	session := s.Copy()
	defer session.Close()

	c := session.DB(banco).C("planets")
	planet.ID = bson.NewObjectId()
	err := c.Insert(&planet)

	return err
}
func (service *planetService) FindAll(s *mgo.Session) []entity.Planet {

	session := s.Copy()
	defer session.Close()

	c := session.DB(banco).C("planets")

	var planet []entity.Planet
	_ = c.Find(bson.M{}).All(&planet)

	return planet
}

func (service *planetService) Find(ps entity.PlanetSearch, s *mgo.Session) []entity.Planet {
	session := s.Copy()
	defer session.Close()

	c := session.DB(banco).C("planets")

	var planetList []entity.Planet

	err := c.Find(ps).All(&planetList)
	checkErr(err)

	return planetList
}

func (service *planetService) MovieCount(ctx *gin.Context, planet entity.Planet) int {

	url := "https://swapi.dev/api/planets/?search=" + planet.Name
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	checkErr(err)

	res, err := client.Do(req)
	checkErr(err)

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	defer res.Body.Close()

	type Objmap struct {
		Count    int         `json:"count"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
		Results  []struct {
			Name           string    `json:"name"`
			RotationPeriod string    `json:"rotation_period"`
			OrbitalPeriod  string    `json:"orbital_period"`
			Diameter       string    `json:"diameter"`
			Climate        string    `json:"climate"`
			Gravity        string    `json:"gravity"`
			Terrain        string    `json:"terrain"`
			SurfaceWater   string    `json:"surface_water"`
			Population     string    `json:"population"`
			Residents      []string  `json:"residents"`
			Films          []string  `json:"films"`
			Created        time.Time `json:"created"`
			Edited         time.Time `json:"edited"`
			URL            string    `json:"url"`
		} `json:"results"`
	}

	var objmap Objmap

	err = json.Unmarshal(body, &objmap)
	checkErr(err)

	ctx.ShouldBindJSON(&objmap)

	for _, values := range objmap.Results {
		return len(values.Films)
	}

	return 0
}

func (service *planetService) Delete(planet entity.PlanetSearch, s *mgo.Session) error {
	session := s.Copy()
	defer session.Close()

	c := session.DB(banco).C("planets")
	err := c.Remove(bson.M{"_id": planet.ID})

	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
