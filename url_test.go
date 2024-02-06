package url

import (
	"fmt"
	"testing"
)

var hostTests = map[string]struct {
	in       string // URL.Host field
	hostname string
	port     string
}{
	"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
	"with empty port": {in: "foo.com:", hostname: "foo.com", port: ""},
	"without port":    {in: "foo.com", hostname: "foo.com", port: ""},
	"IP with port":    {in: "1.2.3.4:90", hostname: "1.2.3.4", port: "90"},
	"IP without port": {in: "1.2.3.4", hostname: "1.2.3.4", port: ""},
}

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"

	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) got %q, want nil", rawurl, err)
	}
	want := "https"
	if got := u.Scheme; got != want {
		t.Errorf("Parse(%q).Scheme got %q, want %q", rawurl, got, want)
	}
	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host got %q, want %q", rawurl, got, want)
	}
	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parse(%q).Path got %q, want %q", rawurl, got, want)
	}
}

func TestURLHost(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("Hostname/%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Hostname(), tt.hostname; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
		t.Run(fmt.Sprintf("Port/%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestURLString(t *testing.T) {
	u := &URL{
		Scheme: "https",
		Host:   "foo.com",
		Path:   "go",
	}

	got, want := u.String(), "https://foo.com/go"
	if got != want {
		t.Errorf("%#v.String()\ngot %q\nwant %q", u, got, want)
	}
}
