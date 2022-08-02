package wowstore

import (
	"testing"
)

func TestPow(t *testing.T) {
	store := NewStore()

	t.Run("Get random quote", func(t *testing.T) {
		str := store.GetRandomQuote()
		if len(str) == 0 {
			t.Errorf("Wrong string: result %d, want >0", len(str))
		}
	})
}
