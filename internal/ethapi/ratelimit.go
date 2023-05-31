package ethapi

import (
	"github.com/ethereum/go-ethereum/common"
)

const RateLimitThreshold = 100 // Default Rate Limit

type RateLimit struct {
	pool map[common.Address]int64
	limitThreshold int64
}

// Get Fresh Rate Limit
func NewRateLimit() *RateLimit {
	rateLimit := &RateLimit{}
	rateLimit.pool = make(map[common.Address]int64)
	rateLimit.limitThreshold = RateLimitThreshold

	return rateLimit
}

// Check whether tx from one signer exceed rate limit, if not, keep tx count in memory
func (r *RateLimit) Add(from common.Address) bool {
	if (r.pool == nil) {
		r.pool = make(map[common.Address]int64)
	}

	value, ok := r.pool[from]

	if (ok) {
		// Check if exceed rate limit
		if (value + 1 > r.limitThreshold) {
			return true
		}

		r.pool[from] = value + 1
	} else {
		r.pool[from] = 1
	}

	return false
}

// Remove tx count once tx is completed
func (r *RateLimit) Remove(from common.Address) {
	value, ok := r.pool[from]

	if (ok && value > 1) {
		r.pool[from] = value - 1
	} else {
		r.pool[from] = 0
	}
}