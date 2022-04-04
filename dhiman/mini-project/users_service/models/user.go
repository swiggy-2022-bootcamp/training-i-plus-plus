package models

// Data Model for a User.
type User struct {
	modelImpl
	Email string `json:"email"` // Email of this client
	Name  string `json:"name"`  // Name of this client
	Password string `json:"password"` // Password of this client
}

// Data Model for a Client.
type Client struct {
	User
	Subscriptions []ClientSubscription `json:"subscriptions"` // Subscriptions that this client has
}

// Generate a new Client with the given data and no subscriptions.
func NewClient(name string, email string, password string) Client {
	c := Client{
		User: User{
			Email: email,
			Name:  name,
			Password: password,
		},
		Subscriptions: []ClientSubscription{},
	}
	c.SetId(email)
	return c
}

// Get the ID of this client.
func (c *Client) GetId() string {
	return c.Email
}

// Data Model for a  Expert.
type Expert struct {
	User
	Specialities []string `json:"specialities"` // Specialities that this Expert has.
}

// Generate a new  Expert with the given data.
func NewExpert(name string, email string, password string, specialities []string) Expert {
	c := Expert{
		User: User{
			Email: email,
			Name:  name,
			Password: password,
		},
		Specialities: specialities,
	}
	c.SetId(email)
	return c
}

// Get the ID of this client.
func (c *Expert) GetId() string {
	return c.Email
}
