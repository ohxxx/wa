// Code generated by "stringer -linecomment -type TLSModel"; DO NOT EDIT.

package llenum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TLSModelNone-0]
	_ = x[TLSModelGeneric-1]
	_ = x[TLSModelInitialExec-2]
	_ = x[TLSModelLocalDynamic-3]
	_ = x[TLSModelLocalExec-4]
}

const _TLSModel_name = "nonegenericinitialexeclocaldynamiclocalexec"

var _TLSModel_index = [...]uint8{0, 4, 11, 22, 34, 43}

func (i TLSModel) String() string {
	if i >= TLSModel(len(_TLSModel_index)-1) {
		return "TLSModel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TLSModel_name[_TLSModel_index[i]:_TLSModel_index[i+1]]
}