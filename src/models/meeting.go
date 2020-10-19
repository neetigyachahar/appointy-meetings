package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meeting struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title"`
	Starttime    string             `json:"starttime"`
	Endtime      string             `json:"endtime"`
	Timestamp    int64              `json:"creation"`
	Participants []*Participants    `json:"participants"`
}
