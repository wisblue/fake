package fake

import (
	"io"

	"crypto/rand"

	"fmt"

	"github.com/segmentio/ksuid"
)

func KSUID() string {
	return ksuid.New().String()
}

// createUUID returns a 16 byte slice with random values
func createUUID() []byte {
	b := make([]byte, 16)
	io.ReadFull(rand.Reader, b)
	// variant bits; see section 4.1.1
	b[8] = b[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	b[6] = b[6]&^0xf0 | 0x40
	return b
}

// Digit returns a 32 bytes UUID
func UUID() string {
	b := createUUID()
	uuid := fmt.Sprintf("%x", b)
	return uuid
}

// OpenID create wechat OpenID string with length of 28
func OpenID() string {
	return text(28, 28, true, true, true, false)
}
