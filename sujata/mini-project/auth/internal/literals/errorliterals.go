package literals

const (
	ParametersMissing     = "all parameters required"
	InvalidEmailFormat    = "email not in format"
	WeakPasswordError     = "password should have atleast 7 length"
	IncorrectUserPassword = "user email or password is incorrect"
	IncorrectUserRole     = "invalid user role please provide correct role"

	DBInsertionError    = "an error occured while inserting in database"
	DBInsertionFail     = "insertion failed as no rows affected "
	DBUserNotFound      = "user not found"
	DBUserAlreadyExists = "user already exists, please sign in"
)
