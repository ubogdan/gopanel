package model

// DnsZone godoc.
type DnsZone struct {
	ID      int64
	Name    string
	Email   string
	TTL     int64
	Refresh int64
	Retry   int64
	Expire  int64
	Minimum int64
	Serial  int64
	Records []DnsRecord
}

// DnsRecord godoc.
type DnsRecord struct {
	ID       int64
	ZoneID   int64
	Zone     DnsZone
	Type     string
	Priority int32
	Weight   int32
	TTL      int64
	Name     string
	Value    string
}
