package ranges

// Ranges are expected to be subsets of one another, and never resulting
// in overlapping ranges
type MatchRange struct {
	Start int
	End   int
}

func (r MatchRange) isSuperset(other MatchRange) bool {
	left := other.Start >= r.Start
	right := r.End >= other.End
	equal := r.Start == other.Start && r.End == other.End

	return left && right && !equal
}

func (r MatchRange) intersects(other MatchRange) bool {
	return r.End >= other.Start && r.Start <= other.End
}

func FindLargest(slices []MatchRange) []MatchRange {
	bestSlices := map[MatchRange]bool{}

	for _, r := range slices {
		if len(bestSlices) == 0 {
			bestSlices[r] = true
			continue
		}

		for bestSlice := range bestSlices {
			if r == bestSlice {
				continue
			} else if r.intersects(bestSlice) {

				// if r ever intersects with another range, delete the subset range.
				if r.isSuperset(bestSlice) {
					delete(bestSlices, bestSlice)
					bestSlices[r] = true
				} else if bestSlices[r] {
					delete(bestSlices, r)
				}
			} else {
				bestSlices[r] = true
			}
		}
	}

	res := []MatchRange{}
	for match := range bestSlices {
		res = append(res, match)
	}

	return res
}
