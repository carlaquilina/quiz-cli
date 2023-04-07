package models

import "sync"

type Results struct {
	resultMutex sync.Mutex
	M           map[string]int
}

func (r *Results) IncreaseScore(username string) {
	r.resultMutex.Lock()
	r.M[username]++
	r.resultMutex.Unlock()
}

func (r Results) GetBetterThan(username string) float64 {
	numUsers := len(r.M)
	if numUsers == 0 {
		return 0
	}

	userScore, ok := r.M[username]
	if !ok {
		return 0
	}

	scoredMore := 0 //score higher than other people
	scoredLess := 0 //scored lower than other people
	for _, score := range r.M {
		if score > userScore {
			scoredLess++
		}
		if score < userScore {
			scoredMore++
		}
	}
	percentBetter := float64(scoredMore) / float64(numUsers-1) * 100
	// percentBetter := float64(rank-1) / float64(numUsers) * 100
	return percentBetter
}
