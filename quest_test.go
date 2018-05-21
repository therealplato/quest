package quest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: fix naming clash between "Success", "Failed", "Complete"
// "In Progress, Failed"

func TestQuest(t *testing.T) {
	q := Q{T: "visit a friend"}
	_ = q
}

func TestLogContainsQuests(t *testing.T) {
	l := Log{
		Quests: []Q{
			Q{T: "pet a dog"},
			Q{T: "pet a bunny"},
		},
	}
	assert.Equal(t, 2, len(l.Quests))
}

// assert.Equal(t, "Accepted Quest", l.history[0])

func TestAchieveQuest(t *testing.T) {
	q := &Q{T: "get 4 viewers on twitch"}
	q.Complete()
	assert.Equal(t, SUCCESS, q.Status)
}

func TestMostRecent(t *testing.T) {
	h := History{"a", "b"}
	assert.Equal(t, "b", h.MostRecent())
}

func TestInProgressQuestCanBeAchieved(t *testing.T) {
	q := &Q{
		T:      "brush teeth",
		Status: PROGRESS,
	}
	q.Complete()
	assert.Equal(t, SUCCESS, q.Status)
}

func TestInProgressQuestCanBeFailed(t *testing.T) {
	q := &Q{
		T:      "brush teeth",
		Status: PROGRESS,
	}
	q.Fail()
	assert.Equal(t, FAILURE, q.Status)
}

func TestFailedQuestCannotBeAchieved(t *testing.T) {
	q := &Q{
		T:      "brush teeth",
		Status: FAILURE,
	}
	q.Complete()
	assert.Equal(t, FAILURE, q.Status)
}

func TestAchievedQuestCannotBeFailed(t *testing.T) {
	q := &Q{
		T:      "brush teeth",
		Status: SUCCESS,
	}
	q.Fail()
	assert.Equal(t, SUCCESS, q.Status)
}
