package ranges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLargestRanges(t *testing.T) {

	t.Run("one large with two subsets", func(t *testing.T) {

		a := MatchRange{Start: 1, End: 4}
		b := MatchRange{Start: 2, End: 5}
		c := MatchRange{Start: 1, End: 10}

		actual := FindLargestRanges([]MatchRange{a, b, c})

		expected := []MatchRange{c}

		assert.Equal(t, expected, actual, "1..10 should be the only range")
	})

	t.Run("two large ranges", func(t *testing.T) {

		a := MatchRange{Start: 1, End: 19}
		b := MatchRange{Start: 4, End: 19}
		c := MatchRange{Start: 21, End: 40}
		d := MatchRange{Start: 25, End: 40}

		actual := FindLargestRanges([]MatchRange{a, b, c, d})

		expected := []MatchRange{a, c}

		assert.ElementsMatch(t, expected, actual, "1..19, 21..40 both ranges")
	})

	t.Run("dupes", func(t *testing.T) {

		a := MatchRange{Start: 1, End: 4}
		b := MatchRange{Start: 1, End: 3}
		c := MatchRange{Start: 1, End: 4}

		actual := FindLargestRanges([]MatchRange{a, b, c})

		expected := []MatchRange{a} // or c

		assert.ElementsMatch(t, expected, actual, "1..4 should be the only range")
	})

	// ranges are expected to be subsets of one another, and never resulting
	// in overlapping ranges
	t.Run("no larger shared superset", func(t *testing.T) {

		a := MatchRange{Start: 1, End: 4}
		b := MatchRange{Start: 2, End: 5}

		actual := FindLargestRanges([]MatchRange{a, b})

		expected := []MatchRange{a, b}

		assert.NotEqual(t, expected, actual, "1..4 should be the only range")
	})
}
