package decoder

import (
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	got, err := Decode(strings.NewReader(jwt))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if want := `{"alg":"HS256","typ":"JWT"}`; want != got.Header {
		t.Errorf("want: %q, got: %q", want, got.Header)
	}
	if want := `{"sub":"1234567890","name":"John Doe","iat":1516239022}`; want != got.Payload {
		t.Errorf("want: %q, got: %q", `{"sub":"1234567890","name":"John Doe","iat":1516239022}`, got.Payload)
	}
}
