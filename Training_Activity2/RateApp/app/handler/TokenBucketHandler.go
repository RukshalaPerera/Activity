package handler

import (
	"RateApp/app/model"
	"math"
	"time"
)

func refill(tb *model.TokenBucket) {
	now := time.Now()
	duration := now.Sub(tb.LastRefill)
	tokensToAdd := tb.RefillRate * duration.Seconds()
	tb.Tokens = math.Min(tb.Tokens+tokensToAdd, tb.MaxTokens)
	tb.LastRefill = now
}

func Request(tb *model.TokenBucket, tokens float64) bool {
	tb.Mu.Lock()
	defer tb.Mu.Unlock()

	refill(tb)
	if tokens <= tb.Tokens {
		tb.Tokens -= tokens
		return true
	}
	return false
}
