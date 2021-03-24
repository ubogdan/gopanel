package model

// Domain godoc.
type Domain struct {
	ID            int64
	Name          string
	DocumentRoot  string
	EnableSSL     bool
	EnablePHP     bool
	EnableCGI     bool
	EnableEmail   bool
	EnableWebMail bool
	PHPVersion    string
}

// Validate godoc.
func (d Domain) Validate() error {
	return nil
}
