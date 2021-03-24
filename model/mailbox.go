package model

// Mailbox godoc.
type Mailbox struct {
	ID             int64
	DomainID       int64
	Domain         *Domain
	Active         bool
	Homedir        string
	Quota          int
	Username       string
	Password       string
	DisablePOP3    bool
	DisableIMAP    bool
	DisableSMTP    bool
	DisableWebmail bool
}

// Mailbox godoc
func (m *Mailbox) Validate() error {
	return nil
}
