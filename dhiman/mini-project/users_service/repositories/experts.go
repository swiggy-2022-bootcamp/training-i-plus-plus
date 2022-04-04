package repositories

import (
	"context"

	"github.com/dhi13man/healthcare-app/users_service/configs"
	"github.com/dhi13man/healthcare-app/users_service/models"
)

// Create a Expert in database.
//
// newExpert models.Expert Expert to insert.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) ID of the created Expert if created, and Error if any occurs.
func CreateExpert(newExpert models.Expert, ctx context.Context) (interface{}, error) {
	res, err := configs.ExpertsCollection.InsertOne(ctx, newExpert)
	return res.InsertedID, err
}

// Get all Experts in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Expert Expert Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  ([]models.Expert, error) Experts if found, and Error if any.
func GetExperts(medicineTemplate models.Expert, ctx context.Context) ([]models.Expert, error) {
	var medicines []models.Expert
	cursor, err := configs.ExpertsCollection.Find(ctx, medicineTemplate)
	if cursor != nil {
		cursor.Decode(&medicines)
	}
	return medicines, err
}


// Get a single Expert in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Expert Expert Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (models.Expert, error) Expert if found, and Error if any.
func GetExpert(medicineTemplate models.Expert, ctx context.Context) (models.Expert, error) {
	var medicine models.Expert
	err := configs.ExpertsCollection.FindOne(ctx, medicineTemplate).Decode(&medicine)
	return medicine, err
}

// Updates Experts in database by updatedExpert's ID (name).
//
// updatedExpert models.Expert Expert to update in database
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) UpsertedID if successful update, and Error if any occurs.
func UpdateExpert(updatedExpert models.Expert, ctx context.Context) (interface{}, error) {
	res, err := configs.ExpertsCollection.UpdateOne(ctx, models.Expert{User: models.User{Name: updatedExpert.Name}}, updatedExpert)
	return res.UpsertedID, err
}

// Delete Experts in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Expert Expert Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (int64, err) The numer of deleted entries, and Error if any occurs.
func DeleteExpert(medicineTemplate models.Expert, ctx context.Context) (int64, error) {
	res, err := configs.ExpertsCollection.DeleteOne(ctx, medicineTemplate)
	return res.DeletedCount, err
}
