package entities

type Client struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewClient(Name string, LastName string, Email string, Password string) *Client {
	return &Client{Name: Name, LastName: LastName, Email: Email, Password: Password}
}
