package url_test

import (
	"fmt"
	"log"
	"pete-dot-m/url"
)

func ExampleURL() {
	u, err := url.Parse("http://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	fmt.Println(u)
	// Output: https://foo.com/go
}
