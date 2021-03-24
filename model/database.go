package model

// Database godoc.
type Database struct {
	ID       int64
	Users    []DatabaseUser
	Name     string
	Type     string
	Port     uint
	Version  string
	User     string
	Password string
}

// DatabaseUser godoc.
type DatabaseUser struct {
	ID         int64
	DatabaseID int64
	Database   *Database
	Username   string
	Password   string
	Rights     []string
}

// Validate godoc.
func (d Database) Validate() error {
	return nil
}

// Validate godoc.
func (d DatabaseUser) Validate() error {
	return nil
}
