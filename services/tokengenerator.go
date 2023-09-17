package services

import (
	"crypto/rand"
	"fmt"
)

func tokenGenerator() string {

	b := make([]byte, 7)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
