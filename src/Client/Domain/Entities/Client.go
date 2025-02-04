package entities

type Client struct {
	ID       int
	Name     string
	LastName string
	Email    string
	Password string
}

func NewClient(Name string, LastName string, Email string, Password string) *Client {
	return &Client{Name: Name, LastName: LastName, Email: Email, Password: Password}
}
