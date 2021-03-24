package model

// UserRights godoc.
type UserRights uint16

// User godoc.
type User struct {
	ID       int64
	Rights   UserRights
	Name     string
	Username string
	Email    string
	Password string
}

// Validate godoc.
func (u User) Validate() error {
	return nil
}
