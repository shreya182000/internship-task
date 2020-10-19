package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meeting struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StartTime         time.Time          `json:"starttime,omitempty" bson:"starttime,omitempty"`
	Title             string             `json:"title" bson:"title,omitempty"`
	Participants      []*Participants    `json:"participants" bson:"participants,omitempty"`
	EndTime           time.Time          `json:"endtime" bson:"endtime,omitempty"`
	CreationTimestamp time.Time          `json:"creationtimestamp" bson:"creationtimestamp,omitempty"`
}
type Participants struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	RSVP  string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}
