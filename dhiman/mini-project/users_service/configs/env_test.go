package configs

import "testing"

// expectedGot is a constant string to format error messages.
const expectedGot string = "Expected a valid string, got %s"

func TestUsersServiceAddress(t *testing.T) {
	// Act
	got := UsersServiceAddress()

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

func TestUsersCollectionNames(t *testing.T) {
	// Act
	got := ClientsCollectionName()
	got2 := ExpertsCollectionName()

	// Assert 
	if got == "" {
		t.Errorf(expectedGot, got)
	}

	// Assert 
	if got2 == "" {
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
