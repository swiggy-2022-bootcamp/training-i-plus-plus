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

func TestUsersCollectionNames(t *testing.T) {
	// Act
	got := ClientsCollectionName()
	got2 := ExpertsCollectionName()

	// Assert 
	if got == "" {
		t.Errorf("Expected a valid string, got %s", got)
	}

	// Assert 
	if got2 == "" {
		t.Errorf("Expected a valid string, got %s", got)
	}
}
