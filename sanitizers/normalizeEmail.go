package sanitizers

import (
	normalizer "github.com/dimuska139/go-email-normalizer/v2"
)

func NormalizeEmail(email string) string {
	n := normalizer.NewNormalizer()
	return n.Normalize(email)
}
