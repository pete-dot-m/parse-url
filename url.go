package url

import (
	"errors"
	"strings"
)

// URL represents the components of a parsed URL string
type URL struct {
	// https://foo.com/go
	Scheme string // https
	Host   string // foo.com
	Path   string // go
}

// Returns the port from the URL, or an empty string
func (u *URL) Port() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return ""
	}
	return u.Host[i+1:]
}

// Returns the hostname
func (u *URL) Hostname() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return u.Host
	}
	return u.Host[:i]
}

// Implementation of the Stringer interface
func (u *URL) String() string {
	if u == nil {
		return ""
	}
	var s strings.Builder
	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		s.WriteString("://")
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		s.WriteString("/")
		s.WriteString(p)
	}
	return s.String()
}

// Parses rawurl into a URL structure
func Parse(rawurl string) (*URL, error) {
	i := strings.Index(rawurl, "://")
	if i < 1 {
		return nil, errors.New("missing scheme")
	}
	scheme, rest := rawurl[:i], rawurl[i+3:]
	host, path := rest, ""
	if i := strings.Index(rest, "/"); i >= 0 {
		host, path = rest[:i], rest[i+1:]
	}
	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil
}
