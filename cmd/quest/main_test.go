package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/therealplato/quest"
)

func TestParseCommand(t *testing.T) {
	// parseCommand("quest", "list")
	// parseCommand("quest", "add")
	// parseCommand("quest", "add", "pay", "bill", "2")
	// parseCommand("quest", "1", "done")
	// parseCommand("quest", "done", "1")
	// $ quest done 1
	// $ quest 1 done
	// $ quest 1 active
	// $ quest 1 lfg
	// $ quest 1 details
	// $ quest 1 restart/reset
	// $ quest 2 details
	// $ quest 999 details
}

func TestInProgressQuestCanBeAchieved(t *testing.T) {
	q1 := &quest.Q{
		T:       "brush teeth",
		Success: false,
		Failed:  false,
		History: nil,
	}
	q1.Achieve()
	// TODO: fix naming clash between "Success", "Failed"
	// "In Progress, Failed"
	assert.True(t, q1.Success)
	assert.False(t, q1.Failed)
}

func TestInProgressQuestCanBeFailed(t *testing.T) {
	q1 := &quest.Q{
		T:       "brush teeth",
		Success: false,
		Failed:  false,
		History: nil,
	}
	q1.Fail()
	assert.False(t, q1.Success)
	assert.True(t, q1.Failed)
}

func TestFailedQuestCannotBeAchieved(t *testing.T) {
	q1 := &quest.Q{
		T:       "brush teeth",
		Success: false,
		Failed:  true,
		History: nil,
	}
	q1.Achieve()
	assert.False(t, q1.Success)
}

func TestAchievedQuestCannotBeFailed(t *testing.T) {
	q1 := &quest.Q{
		T:       "brush teeth",
		Success: true,
		Failed:  false,
		History: nil,
	}
	q1.Fail()
	assert.True(t, q1.Success)
}
