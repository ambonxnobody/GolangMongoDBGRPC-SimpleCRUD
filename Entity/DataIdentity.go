package Entity

type DataIdentityEntity struct {
	CreatedAt int64 `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	UpdatedAt int64 `json:"UpdatedAt,omitempty" bson:"UpdatedAt,omitempty"`
	DeletedAt int64 `json:"DeletedAt,omitempty" bson:"DeletedAt,omitempty"`

	CreatedBy string `json:"CreatedBy,omitempty" bson:"CreatedBy,omitempty"`
	UpdatedBy string `json:"UpdatedBy,omitempty" bson:"UpdatedBy,omitempty"`
	DeletedBy string `json:"DeletedBy,omitempty" bson:"DeletedBy,omitempty"`
}
