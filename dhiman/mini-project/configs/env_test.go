package configs

import "testing"

func TestEnvMongoURI(t *testing.T) {
	// Act
	got := EnvMongoURI()

	// Assert 
	if got == "" {
		t.Errorf("Expected a valid string, got %s", got)
	}
}

func TestUsersCollectionName(t *testing.T) {
	// Act
	got := UsersCollectionName()

	// Assert 
	if got == "" {
		t.Errorf("Expected a valid string, got %s", got)
	}
}

func TestMedicinesCollectionName(t *testing.T) {
	// Act
	got := MedicinesCollectionName()

	// Assert 
	if got == "" {
		t.Errorf("Expected a valid string, got %s", got)
	}
}
