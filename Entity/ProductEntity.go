package Entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductEntity struct {
	ID         primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name       string             `json:"Name" bson:"Name"`
	Price      int64              `json:"Price" bson:"Price"`
	Stock      int64              `json:"Stock" bson:"Stock"`
	CategoryID string             `json:"CategoryID" bson:"CategoryID"`
}
