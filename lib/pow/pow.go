package pow

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"wow/config"
)

type Pow struct {
	Nonce int
	Hash  string
}

var difficulty int

func init() {
	rand.Seed(time.Now().Unix())
	_, _, difficulty = config.GetConfig()
}

func New() Pow {
	nonce := rand.Intn(difficulty/2) + difficulty/2

	secret := sha1.New()
	secret.Write([]byte(fmt.Sprintf("%d", nonce)))
	secretHash := fmt.Sprintf("%x", secret.Sum(nil))
	return Pow{
		Nonce: nonce,
		Hash:  secretHash,
	}
}

func (p *Pow) Update() {
	nonce := rand.Intn(difficulty/2) + difficulty/2

	secret := sha1.New()
	secret.Write([]byte(fmt.Sprintf("%d", nonce)))
	secretHash := fmt.Sprintf("%x", secret.Sum(nil))

	p.Nonce = nonce
	p.Hash = secretHash
}

func (p *Pow) Solve() error {
	for i := 0; i < difficulty; i++ {
		secret := sha1.New()
		secret.Write([]byte(fmt.Sprintf("%d", i)))
		hash := fmt.Sprintf("%x", secret.Sum(nil))
		if hash == p.Hash {
			p.Nonce = i
			p.Hash = hash
			return nil
		}
	}

	return errors.New("Failed to find challenge answer")
}
