package utils

import "crypto/rand"

func GenerateId(prefix string) string {
	return prefix + "_" + rand.Text()[:9]
}
