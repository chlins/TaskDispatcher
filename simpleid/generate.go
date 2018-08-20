package simpleid

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

var (
	errIllegalArguments = errors.New("invaild argument")
)

var symbols = "-_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// New generates a simple id by default length.
func New() (string, error) {
	// default length is 8
	return NewWithLength(8)
}

// NewWithLength generates a simple id by specified length.
func NewWithLength(length int) (string, error) {
	if length < 0 {
		return "", errIllegalArguments
	}

	// set random seeds
	rand.Seed(time.Now().UnixNano())

	// generate id
	var res strings.Builder
	for i := 0; i < length; i++ {
		index := rand.Int31n(int32(len(symbols)))
		err := res.WriteByte(symbols[index])
		if err != nil {
			return "", err
		}
	}

	return res.String(), nil
}
