package rand

import (
	"encoding/base64"
	"fmt"

	"crypto/rand"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if nRead < n {
		return nil, fmt.Errorf("bytes: didn't read enough random bytes")
	}
	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}
	return b, nil
}

// String returns a random string using crypto/rand.
// n is the number of bytes used to generate the random string.
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
