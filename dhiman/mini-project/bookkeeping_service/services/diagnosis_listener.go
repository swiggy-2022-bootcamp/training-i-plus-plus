package services

import (
	"context"
	"encoding/json"

	"github.com/dhi13man/healthcare-app/bookkeeping_service/models"
	"github.com/dhi13man/healthcare-app/bookkeeping_service/repositories"
	log "github.com/sirupsen/logrus"
)

// Takes a serialized Disease diagnosis string and deserializes it into a Disease object.
//
// Then saves the Disease to the database.
func DeserializeAndSaveDiseaseDiagnosis(serialDisease string, ctx context.Context) {
	var disease models.Disease
	err := json.Unmarshal([]byte(serialDisease), &disease)
	if err != nil {
		log.Error("Error deserializing disease: ", err)
	} else {
		out, err := repositories.CreateDisease(disease, ctx)
		if err != nil {
			log.Error("Error saving disease: ", err)
		} else {
			log.Info("Saved disease: ", out)
		}
	}
}