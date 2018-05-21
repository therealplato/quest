package quest

import "github.com/therealplato/subverse"

// Q is a quest
type Q struct {
	T       string
	History History
}

// Log is a hero's quest list
type Log struct {
	Hero   subverse.Identity
	Quests []Q
}

// History describes quest progress
type History []string
