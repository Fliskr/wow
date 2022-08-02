package wowstore

import (
	"math/rand"
	"time"
)

var initialData = []string{
	"Self-Improvement and success go hand in hand. Taking the steps to make yourself a better and more well-rounded individual will prove to be a wise decision.",
	"The wise person feels the pain of one arrow. The unwise feels the pain of two",
	"When looking for wise words, the best ones often come from our elders",
	"You've heard that it's wise to learn from experience, but it is wiser to learn from the experience of others.",
	"We tend to think of great thinkers and innovators as soloists, but the truth is that the greatest innovative thinking doesn't occur in a vacuum. Innovation results from collaboration.",
}

func init() {
	rand.Seed(time.Now().Unix())
}

type storeImpl struct{}

var store *storeImpl

func NewStore() *storeImpl {
	if store == nil {
		store = &storeImpl{}
	}
	return store
}

func (s *storeImpl) GetRandomQuote() string {
	return initialData[rand.Intn(len(initialData))]
}
