package model

type User struct {
	Name       string `bson:"name" json:"name"`
	ProfileUrl string `bson:"profileUrl" json:"profileUrl"`
}
