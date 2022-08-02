package pow

import (
	"testing"
	"wow/config"
)

func TestPow(t *testing.T) {
	p := New()
	p.Nonce = 0

	nonce := p.Nonce
	p.Update()
	t.Run("Pow updated", func(t *testing.T) {
		if nonce == p.Nonce {
			t.Errorf("Pow wasn't updated: result = 0, want = random [%d-%d)", config.DIFFICULTY/2, config.DIFFICULTY)
		}
	})

	t.Run("Pow solved", func(t *testing.T) {
		p.Nonce = 0
		p.Hash = "d70283dd9dc9ab87833fa5b5e2ebed8a5850b896"

		err := p.Solve()
		if err != nil {
			t.Error("Error solving pow")
		}

		if p.Nonce != 1072121 {
			t.Errorf("data:%s %d", p.Hash, p.Nonce)
		}

	})
}
