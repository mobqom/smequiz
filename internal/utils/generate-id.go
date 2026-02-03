package utils

import "crypto/rand"

func GenerateId(prefix string) string {
	return rand.Text()[:9]
}
