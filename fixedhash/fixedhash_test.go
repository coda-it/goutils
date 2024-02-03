package fixedhash

import (
	"testing"
)

func TestFixedHash(t *testing.T) {
	t.Run("Should return error when `length` is not multiple of 4", func(t *testing.T) {
		hash, err := FixedHash(3)

		if hash == "" && err == nil {
			t.Errorf("Function didn't return error for incorrect input")
		}
	})

	t.Run("Should generate unique hash every time function is called", func(t *testing.T) {
		hash1, err := FixedHash(4)

		if len(hash1) != 4 && err != nil {
			t.Errorf("Hash length should be 4 but is %v", len(hash1))
		}

		hash2, err := FixedHash(8)

		if len(hash2) != 8 && err != nil {
			t.Errorf("Hash length should be 3 but is %v", len(hash2))
		}

		if hash1 == hash2 {
			t.Errorf("Hash 1 should be different than hash 2")
		}
	})
}
