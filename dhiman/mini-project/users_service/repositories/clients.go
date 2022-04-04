package repositories

import (
	"context"

	"github.com/dhi13man/healthcare-app/users_service/configs"
	"github.com/dhi13man/healthcare-app/users_service/models"
)

// Create a Client in database.
//
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) ID of the created Client if created, and Error if any occurs.
func CreateClient(newClient models.Client, ctx context.Context) (interface{}, error) {
	res, err := configs.ClientsCollection.InsertOne(ctx, newClient)
	return res.InsertedID, err
}

// Get all Clients in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Client Client Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  ([]models.Client, error) Clients if found, and Error if any.
func GetClients(medicineTemplate models.Client, ctx context.Context) ([]models.Client, error) {
	var medicines []models.Client
	cursor, err := configs.ClientsCollection.Find(ctx, medicineTemplate)
	if cursor != nil {
		cursor.Decode(&medicines)
	}
	return medicines, err
}


// Get a single Client in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Client Client Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (models.Client, error) Client if found, and Error if any.
func GetClient(medicineTemplate models.Client, ctx context.Context) (models.Client, error) {
	var medicine models.Client
	err := configs.ClientsCollection.FindOne(ctx, medicineTemplate).Decode(&medicine)
	return medicine, err
}

// Updates Clients in database by updatedClient's ID (name).
//
// updatedClient models.Client Client to update in database
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) UpsertedID if successful update, and Error if any occurs.
func UpdateClient(updatedClient models.Client, ctx context.Context) (interface{}, error) {
	res, err := configs.ClientsCollection.UpdateOne(ctx, models.Client{User: models.User{Name: updatedClient.Name}}, updatedClient)
	return res.UpsertedID, err
}

// Delete Clients in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Client Client Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (int64, err) The numer of deleted entries, and Error if any occurs.
func DeleteClient(medicineTemplate models.Client, ctx context.Context) (int64, error) {
	res, err := configs.ClientsCollection.DeleteOne(ctx, medicineTemplate)
	return res.DeletedCount, err
}
