package v0

type Printer struct {
	host    string
	key     string
	headers map[string]string
}

func NewPrinter(host string, key string) *Printer {
	return &Printer{
		host: host,
		key:  key,
		headers: map[string]string{
			"X-Api-Key": key,
		},
	}
}

func (p Printer) Version() (*Version, error) {
	version := &Version{}
	return getAndParseAsJSON(p, "/api/version", version)
}
