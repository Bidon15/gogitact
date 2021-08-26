package computation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCalculator(t *testing.T) {
	le_meme := New("lol kek cheburek")
	num, _ := le_meme.GetWiser()
	assert.GreaterOrEqual(t, num, 0)
}
