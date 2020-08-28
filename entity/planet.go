package entity

import "gopkg.in/mgo.v2/bson"

type Planet struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name      string        `json:"name,omitempty" binding:"min=4,max=25" `
	Climate   string        `json:"climate" binding:"min=3,max=25`
	Terrain   string        `json:"terrain" binding:"min=3,max=25`
	FilmCount int           `json:"film_count`
}
type PlanetSearch struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name,omitempty" json:"name"`
}
