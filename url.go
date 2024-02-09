package url

import (
	"errors"
	"fmt"
	"strings"
)

// URL represents the components of a parsed URL string
type URL struct {
	// https://foo.com/go
	Scheme string // https
	Host   string // foo.com
	Path   string // go
}

// Parses rawurl into a URL structure
func Parse(rawurl string) (*URL, error) {
	scheme, rest, ok := parseScheme(rawurl)
	if !ok {
		return nil, errors.New("missing scheme")
	}

	host, path := parseHostPath(rest)

	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil
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

func (u *URL) testString() string {
	return fmt.Sprintf("scheme=%q, host=%q, path=%q", u.Scheme, u.Host, u.Path)
}

func parseScheme(rawurl string) (scheme, rest string, ok bool) {
	return split(rawurl, "://", 1)
}

func parseHostPath(hostpath string) (host, path string) {
	host, path, ok := split(hostpath, "/", 0)
	if !ok {
		host = hostpath
	}
	return host, path
}

func split(s, sep string, n int) (a, b string, ok bool) {
	i := strings.Index(s, sep)
	if i < n {
		return "", "", false
	}
	return s[:i], s[i+len(sep):], true
}
