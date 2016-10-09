package model

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Presentation struct {
	Id        string    `bson:"id" json:"id"`
	Title     string    `bson:"title" json:"title"`
	ThumbNail string    `bson:"thumbnail" json:"thumbNail"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

/*This is to handle the custom time format*/
func (presentation *Presentation) SetBSON(raw bson.Raw) error {
	result := new(struct {
		Id        string `bson:"id" json:"id"`
		Title     string `bson:"title" json:"title"`
		ThumbNail string `bson:"thumbnail" json:"thumbNail"`
		CreatedAt string `bson:"createdAt" json:"createdAt"`
	})

	err := raw.Unmarshal(result)

	if err != nil {
		return err
	}

	presentation.Id = result.Id
	presentation.Title = result.Title
	presentation.ThumbNail = result.ThumbNail

	form := "January _2, 2006"
	t, e := time.Parse(form, result.CreatedAt)

	if e != nil {
		log.Println(e.Error())
	}

	presentation.CreatedAt = t

	return nil
}
