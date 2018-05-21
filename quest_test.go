package quest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	q.Achieve()
	assert.True(t, q.Success)
}

func TestMostRecent(t *testing.T) {
	h := History{"a", "b"}
	assert.Equal(t, "b", h.MostRecent())
}
