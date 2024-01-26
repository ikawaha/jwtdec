package decoder

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type JWTDecoded struct {
	Header  string
	Payload string
}

func Decode(r io.Reader) (*JWTDecoded, error) {
	jwt, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}
	s := strings.Split(string(jwt), ".")
	if len(s) != 3 {
		return nil, fmt.Errorf("invalid jwt format")
	}
	var ret JWTDecoded
	if b, err := base64.RawURLEncoding.DecodeString(s[0]); err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	} else if !json.Valid(b) {
		return nil, fmt.Errorf("invalid json header %q", s[0])
	} else {
		ret.Header = string(b)
	}
	if b, err := base64.RawURLEncoding.DecodeString(s[1]); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	} else if !json.Valid(b) {
		return nil, fmt.Errorf("invalid json payload %q", s[1])
	} else {
		ret.Payload = string(b)
	}
	return &ret, nil
}
