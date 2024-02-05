package url

import "testing"

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
