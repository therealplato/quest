package quest

import (
	"fmt"

	"github.com/therealplato/subverse"
)

// Q is a quest
type Q struct {
	T       string
	Done    bool
	History History
}

func (q *Q) Achieve() {
	fmt.Println("achieving")
	q.Done = true
}

// Log is a hero's quest list
type Log struct {
	Hero   subverse.Identity
	Quests []Q
}

// History describes quest progress
type History []string
