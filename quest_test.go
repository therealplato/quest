package quest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// define: democracy, autocracy, voluntarism
// visit a friend
// brush teeth
// find information you seek
// share information with someone who doesn't know they seek it

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
