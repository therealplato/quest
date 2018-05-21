package quest

import (
	"github.com/therealplato/subverse"
)

const (
	FAILURE  Status = -1
	PROGRESS Status = 0
	SUCCESS  Status = 1
)

// Q is a quest
type Q struct {
	// T is the text of the quest
	T string
	// Status is -1 if quest failed, 0 in progress, 1 success
	Status Status
	// History is what progress you've made
	History History
}

type Status int

func (s Status) String() string {
	switch s {
	case FAILURE:
		return "X"
	case PROGRESS:
		return " "
	case SUCCESS:
		// return "✓"
		return "+"
	}
	return ""
}

// Achieve marks the quest as done
func (q *Q) Achieve() {
	if q.Status != PROGRESS {
		return
	}
	q.Status = SUCCESS
	q.History = append(q.History, "Quest Achieved")
}

// Fail marks the quest as permanently failed
func (q *Q) Fail() {
	if q.Status != PROGRESS {
		return
	}
	q.Status = FAILURE
}

// Log is a hero's quest list
type Log struct {
	Hero   subverse.Identity
	Quests []Q
}

// History describes quest progress
type History []string

func (h History) MostRecent() string {
	if len(h) == 0 {
		return ""
	}
	return h[len(h)-1]
}
