package api

// User represents a user in the system.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUser returns a sample user.
func GetUser() User {
	return User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
}
