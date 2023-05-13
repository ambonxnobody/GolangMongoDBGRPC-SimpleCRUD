package Entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryEntity struct {
	ID   primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name string             `json:"Name" bson:"Name"`

	DataIdentity DataIdentityEntity `json:"DataIdentity,omitempty" bson:"DataIdentity,omitempty"`
}
