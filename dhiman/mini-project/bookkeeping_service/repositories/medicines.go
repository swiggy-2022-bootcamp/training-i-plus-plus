package repositories

import (
	"context"

	"github.com/dhi13man/healthcare-app/bookkeeping_service/configs"
	"github.com/dhi13man/healthcare-app/bookkeeping_service/models"
)

/**
 * Create a Medicine in database.
 * @param newMedicine models.Medicine Medicine to insert.
 * @param c *context.Context Context to control deadline, cancellation signal, etc.
 * @return (interface{}, error) ID of the created Medicine if created, and Error if any occurs.
 */
func CreateMedicine(newMedicine models.Medicine, ctx context.Context) (interface{}, error) {
	res, err := configs.MedicinesCollection.InsertOne(ctx, newMedicine)
	return res.InsertedID, err
}

/**
 * Get all Medicines in database where fields match medicineTemplate filter.
 * @param medicineTemplate models.Medicine Medicine Template to filter data by.
 * @param c *context.Context Context to control deadline, cancellation signal, etc.
 * @return ([]models.Medicine, error) Medicines if found, and Error if any.
 */
func GetMedicines(medicineTemplate models.Medicine, ctx context.Context) ([]models.Medicine, error) {
	var medicines []models.Medicine
	cursor, err := configs.MedicinesCollection.Find(ctx, medicineTemplate)
	if cursor != nil {
		cursor.Decode(&medicines)
	}
	return medicines, err
}

/**
 * Get a single Medicine in database where fields match medicineTemplate filter.
 * @param medicineTemplate models.Medicine Medicine Template to filter data by.
 * @param c *context.Context Context to control deadline, cancellation signal, etc.
 * @return (models.Medicine, error) The Medicine if found, and Error if any occurs.
 */
func GetMedicine(medicineTemplate models.Medicine, ctx context.Context) (models.Medicine, error) {
	var medicine models.Medicine
	err := configs.MedicinesCollection.FindOne(ctx, medicineTemplate).Decode(&medicine)
	return medicine, err
}

/**
 * Updates Medicines in database by updatedMedicine's ID (name).
 * @param updatedMedicine models.Medicine Medicine to update in database
 * @param c *context.Context Context to control deadline, cancellation signal, etc.
 * @return (interface{}, error) UpsertedID if successful update, and Error if any occurs.
 */
func UpdateMedicine(updatedMedicine models.Medicine, ctx context.Context) (interface{}, error) {
	res, err := configs.MedicinesCollection.UpdateOne(ctx, models.Medicine{Name: updatedMedicine.Name}, updatedMedicine)
	return res.UpsertedID, err
}

/**
 * Delete Medicines in database where fields match medicineTemplate filter.
 * @param medicineTemplate models.Medicine Medicine Template to filter data by.
 * @param c *context.Context Context to control deadline, cancellation signal, etc.
 * @return (int64, err) The numer of deleted entries, and Error if any occurs.
 */
func DeleteMedicine(medicineTemplate models.Medicine, ctx context.Context) (int64, error) {
	res, err := configs.MedicinesCollection.DeleteOne(ctx, medicineTemplate)
	return res.DeletedCount, err
}
