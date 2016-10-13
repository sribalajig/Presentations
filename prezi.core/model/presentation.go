package model

import (
	"time"
)

type Presentation struct {
	Id        string    `bson:"id" json:"id"`
	Title     string    `bson:"title" json:"title"`
	ThumbNail string    `bson:"thumbnail" json:"thumbNail"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	Creator   User      `bson:"creator" json:"creator"`
}
