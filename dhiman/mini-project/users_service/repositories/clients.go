package repositories

import (
	"context"

	"github.com/dhi13man/healthcare-app/users_service/configs"
	"github.com/dhi13man/healthcare-app/users_service/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Create a Client in database.
//
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) ID of the created Client if created, and Error if any occurs.
func CreateClient(newClient models.Client, ctx context.Context) (interface{}, error) {
	res, err := configs.ClientsCollection.InsertOne(ctx, newClient)
	if res == nil {
		return nil, err
	} else {
		return res.InsertedID, err
	}
}

// Get all Clients in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Client Client Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  ([]models.Client, error) Clients if found, and Error if any.
func GetClients(userTemplate models.Client, ctx context.Context) ([]models.Client, error) {
	var users []models.Client
	cursor, err := configs.ClientsCollection.Find(ctx, userTemplate)
	if cursor != nil {
		cursor.Decode(&users)
	}
	return users, err
}


// Get a single Client in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Client Client Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (models.Client, error) Client if found, and Error if any.
func GetClient(userTemplate models.Client, ctx context.Context) (models.Client, error) {
	var client models.Client
	err := configs.ClientsCollection.FindOne(ctx, bson.M{
		"email": userTemplate.Email,
	}).Decode(&client)
	return client, err
}

// Updates Clients in database by updatedClient's ID (email).
//
// updatedClient models.Client Client to update in database
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (interface{}, error) UpsertedID if successful update, and Error if any occurs.
func UpdateClient(updatedClient models.Client, ctx context.Context) (interface{}, error) {
	res, err := configs.ClientsCollection.UpdateOne(
		ctx, 
		bson.M{
			"email": updatedClient.Email,
		}, 
		bson.M{
			"$set": updatedClient,
		},
	)
	if res == nil {
		return nil, err
	} else {
		return res.UpsertedID, err
	}
}

// Delete Clients in database where fields match medicineTemplate filter.
//
// medicineTemplate models.Client Client Template to filter data by.
// c *context.Context Context to control deadline, cancellation signal, etc.
//
// @return  (int64, err) The numer of deleted entries, and Error if any occurs.
func DeleteClient(medicineTemplate models.Client, ctx context.Context) (int64, error) {
	res, err := configs.ClientsCollection.DeleteOne(ctx,  bson.M{
		"email": medicineTemplate.Email,
	})
	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
