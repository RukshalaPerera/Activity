package model

import (
	"sync"
	"time"
)

type TokenBucket struct {
	Tokens     float64
	MaxTokens  float64
	RefillRate float64
	LastRefill time.Time
	Mu         sync.Mutex
}

func NewTokenBucket(maxTokens float64, refillRate float64) *TokenBucket {
	return &TokenBucket{
		Tokens:     maxTokens,
		MaxTokens:  maxTokens,
		RefillRate: refillRate,
		LastRefill: time.Now(),
	}
}
