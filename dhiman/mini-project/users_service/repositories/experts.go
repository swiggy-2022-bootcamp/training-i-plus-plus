package repositories

import (
	"context"

	"github.com/dhi13man/healthcare-app/users_service/configs"
	"github.com/dhi13man/healthcare-app/users_service/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Create a Expert in database.
//
// newExpert models.Expert Expert to insert.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) ID of the created Expert if created, and Error if any occurs.
func CreateExpert(newExpert models.Expert, ctx context.Context) (interface{}, error) {
	res, err := configs.ExpertsCollection.InsertOne(ctx, newExpert)
	if res == nil {
		return nil, err
	} else {
		return res.InsertedID, err
	}
}

// Get all Experts in database where fields match expertTemplate filter.
//
// expertTemplate models.Expert Expert Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  ([]models.Expert, error) Experts if found, and Error if any.
func GetExperts(expertTemplate models.Expert, ctx context.Context) ([]models.Expert, error) {
	var medicines []models.Expert
	cursor, err := configs.ExpertsCollection.Find(ctx, bson.M{"email": expertTemplate.Email})
	if cursor != nil {
		cursor.Decode(&medicines)
	}
	return medicines, err
}


// Get a single Expert in database where fields match expertTemplate filter.
//
// expertTemplate models.Expert Expert Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (models.Expert, error) Expert if found, and Error if any.
func GetExpert(expertTemplate models.Expert, ctx context.Context) (models.Expert, error) {
	var medicine models.Expert
	err := configs.ExpertsCollection.FindOne(ctx, bson.M{"email": expertTemplate.Email}).Decode(&medicine)
	return medicine, err
}

// Updates Experts in database by updatedExpert's ID (name).
//
// updatedExpert models.Expert Expert to update in database
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) UpsertedID if successful update, and Error if any occurs.
func UpdateExpert(updatedExpert models.Expert, ctx context.Context) (interface{}, error) {
	res, err := configs.ExpertsCollection.UpdateOne(ctx, bson.M{"email": updatedExpert.Email}, updatedExpert)
	if res == nil {
		return nil, err
	} else {
		return res.UpsertedID, err
	}
}

// Delete Experts in database where fields match expertTemplate filter.
//
// expertTemplate models.Expert Expert Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (int64, err) The numer of deleted entries, and Error if any occurs.
func DeleteExpert(expertTemplate models.Expert, ctx context.Context) (int64, error) {
	res, err := configs.ExpertsCollection.DeleteOne(ctx, bson.M{"email": expertTemplate.Email})
	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
