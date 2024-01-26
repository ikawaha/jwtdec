Tiny JWT Decoder
================

This is a tiny JWT decoder. It just displays the header and payload from the jwt string.

## Installation

```bash
go install github.com/ikawaha/jwtdec@latest
```

## Usage

```bash
jwtdec <jwt>
```

```bash
echo <jwt> | jwtdec 
```
---

### Example

```bash
jwtdec eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ. \
 | jq .
```
```ndjson
{
  "alg": "HS256",
  "typ": "JWT"
}
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```
---
MIT