package entity

type Space struct {
	SpaceID   string `bson:"space_id" json:"space_id"`
	Name      string `bson:"name" json:"name"`
	Location  string `bson:"location" json:"location"`
	Capacity  int32  `bson:"capacity" json:"capacity"`
	Available bool   `bson:"available" json:"available"`
}
