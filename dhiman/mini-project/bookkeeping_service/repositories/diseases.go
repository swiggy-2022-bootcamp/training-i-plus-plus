package repositories

import (
	"context"

	"github.com/dhi13man/healthcare-app/bookkeeping_service/configs"
	"github.com/dhi13man/healthcare-app/bookkeeping_service/models"
)

// Create a Disease in database.
//
// newDisease models.Disease Disease to insert.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) ID of the created Disease if created, and Error if any occurs.
func CreateDisease(newDisease models.Disease, ctx context.Context) (interface{}, error) {
	res, err := configs.DiseasesCollection.InsertOne(ctx, newDisease)
	if res == nil {
		return nil, err
	} else {
		return res.InsertedID, err
	}
}

// Get all Diseases in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Disease Disease Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  ([]models.Disease, error) Diseases if found, and Error if any.
func GetDiseases(medicineTemplate models.Disease, ctx context.Context) ([]models.Disease, error) {
	var medicines []models.Disease
	cursor, err := configs.DiseasesCollection.Find(ctx, medicineTemplate)
	if cursor != nil {
		cursor.Decode(&medicines)
	}
	return medicines, err
}

// Get a single Disease in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Disease Disease Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (models.Disease, error) The Disease if found, and Error if any occurs.
func GetDisease(medicineTemplate models.Disease, ctx context.Context) (models.Disease, error) {
	var medicine models.Disease
	err := configs.DiseasesCollection.FindOne(ctx, medicineTemplate).Decode(&medicine)
	return medicine, err
}

// Updates Diseases in database by updatedDisease's ID (name).
//
// updatedDisease models.Disease Disease to update in database
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) UpsertedID if successful update, and Error if any occurs.
func UpdateDisease(updatedDisease models.Disease, ctx context.Context) (interface{}, error) {
	res, err := configs.DiseasesCollection.UpdateOne(ctx, models.Disease{Name: updatedDisease.Name}, updatedDisease)
	if res == nil {
		return nil, err
	} else {
		return res.UpsertedID, err
	}
}

// Delete Diseases in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Disease Disease Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (int64, err) The numer of deleted entries, and Error if any occurs.
func DeleteDisease(medicineTemplate models.Disease, ctx context.Context) (int64, error) {
	res, err := configs.DiseasesCollection.DeleteOne(ctx, medicineTemplate)
	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
