package domain

type Space struct {
	ID        string `bson:"_id"`
	Name      string
	Available int
}

type SpaceRepository interface {
	DecreaseAvailability(spaceID string) error
}
