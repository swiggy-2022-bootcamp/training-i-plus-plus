package models


// Data Model for a Client.
type Client struct {
	Email string `json:"email"` // Email of this client
	Name  string `json:"name"`  // Name of this client
	Password string `json:"password"` // Password of this client
	Subscriptions []ClientSubscription `json:"subscriptions"` // Subscriptions that this client has
}

// Generate a new Client with the given data and no subscriptions.
func NewClient(name string, email string, password string, subscriptions []ClientSubscription) Client {
	c := Client{
		Email: email,
		Name:  name,
		Password: password,
		Subscriptions: subscriptions,
	}
	return c
}

// Get the ID of this client.
func (c *Client) GetId() string {
	return c.Email
}

// Data Model for a  Expert.
type Expert struct {
	Email string `json:"email"` // Email of this client
	Name  string `json:"name"`  // Name of this client
	Password string `json:"password"` // Password of this client
	Specialities []string `json:"specialities"` // Specialities that this Expert has.
}

// Generate a new  Expert with the given data.
func NewExpert(name string, email string, password string, specialities []string) Expert {
	c := Expert{
		Email: email,
		Name:  name,
		Password: password,
		Specialities: specialities,
	}
	return c
}

// Get the ID of this client.
func (c *Expert) GetId() string {
	return c.Email
}
