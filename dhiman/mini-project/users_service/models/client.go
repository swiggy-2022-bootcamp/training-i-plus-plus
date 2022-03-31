package models

type Client struct {
	modelImpl
	Email         string         `json:"email"`         // Email of this client
	Name          string         `json:"name"`          // Name of this client
	Subscriptions []UserSubscription `json:"subscriptions"` // Subscriptions that this client has
}

func NewClient(name string, email string) *Client {
	c := &Client{
		Email:         email,
		Name:          name,
		Subscriptions: []UserSubscription{},
	}
	c.SetId(email)
	return c
}

func (c *Client) GetId() string {
	return c.Email
}
