package fixedhash

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math"
)

func FixedHash(length int) (string, error) {
	if length%4 != 0 {
		return "", errors.New("length should be multiple of 4")
	}

	sourceLength := int(math.Ceil((float64(length) * 3) / 4))

	var hashSource = make([]byte, sourceLength)

	_, err := rand.Read(hashSource)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(hashSource), nil
}
