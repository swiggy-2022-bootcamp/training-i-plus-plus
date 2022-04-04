package configs

import "testing"

// expectedGot is a constant string to format error messages.
const expectedGot string = "Expected a valid string, got %s"

func TestBookkeepingServiceAddress(t *testing.T) {
	// Act
	got := BookkeepingServiceAddress()

	// Assert 
	if got == "" {
		t.Errorf(expectedGot, got)
	}
}

func TestEnvMongoURI(t *testing.T) {
	// Act
	got := EnvMongoURI()

	// Assert 
	if got == "" {
		t.Errorf(expectedGot, got)
	}
}

func TestDiseasesCollectionName(t *testing.T) {
	// Act
	got := DiseasesCollectionName()

	// Assert 
	if got == "" {
		t.Errorf(expectedGot, got)
	}
}

func TestMedicinesCollectionName(t *testing.T) {
	// Act
	got := MedicinesCollectionName()

	// Assert 
	if got == "" {
		t.Errorf(expectedGot, got)
	}
}

func TestKafkaBrokerAddress(t *testing.T) {
	// Act
	got := KafkaBrokerAddress()

	// Assert 
	if got == "" {
		t.Errorf(expectedGot, got)
	}
}

func TestKafkaDiagnosisTopic(t *testing.T) {
	// Act
	got := KafkaDiagnosisTopic()

	// Assert 
	if got == "" {
		t.Errorf(expectedGot, got)
	}
}
