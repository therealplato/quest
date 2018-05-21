package quest

import (
	"fmt"

	"github.com/therealplato/subverse"
)

// Q is a quest
type Q struct {
	T       string
	Success bool
	Failed  bool
	History History
}

func (q *Q) Achieve() {
	if q.Failed {
		return
	}
	fmt.Printf("Quest Achieved: %s\n", q.T)
	q.Success = true
}
func (q *Q) Fail() {
	if q.Success {
		return
	}
	fmt.Printf("Quest Failed: %s\n", q.T)
	q.Success = false
	q.Failed = false
}

// Log is a hero's quest list
type Log struct {
	Hero   subverse.Identity
	Quests []Q
}

// History describes quest progress
type History []string
