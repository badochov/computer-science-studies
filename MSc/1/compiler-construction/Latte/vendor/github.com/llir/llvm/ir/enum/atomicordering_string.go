// Code generated by "stringer -linecomment -type AtomicOrdering"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AtomicOrderingNone-0]
	_ = x[AtomicOrderingAcqRel-1]
	_ = x[AtomicOrderingAcquire-2]
	_ = x[AtomicOrderingMonotonic-3]
	_ = x[AtomicOrderingRelease-4]
	_ = x[AtomicOrderingSeqCst-5]
	_ = x[AtomicOrderingUnordered-6]
}

const _AtomicOrdering_name = "noneacq_relacquiremonotonicreleaseseq_cstunordered"

var _AtomicOrdering_index = [...]uint8{0, 4, 11, 18, 27, 34, 41, 50}

func (i AtomicOrdering) String() string {
	if i >= AtomicOrdering(len(_AtomicOrdering_index)-1) {
		return "AtomicOrdering(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AtomicOrdering_name[_AtomicOrdering_index[i]:_AtomicOrdering_index[i+1]]
}