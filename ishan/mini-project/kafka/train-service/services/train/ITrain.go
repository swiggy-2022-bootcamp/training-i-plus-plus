package train

import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	ID          primitive.ObjectID `bson:"_id",omitempty`
	Name        string             `bson:"name"`
	Number      string             `bson:"number"`
	Destination string             `bson:"destination"`
	Source      string             `bson:"source"`
	Stations    []string           `bson:"stations"`
	Seats       int32              `bson:"seats"`
}
