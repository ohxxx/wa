// Code generated by "stringer -linecomment -type ReturnAttr"; DO NOT EDIT.

package llenum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ReturnAttrInReg-0]
	_ = x[ReturnAttrNoAlias-1]
	_ = x[ReturnAttrNoMerge-2]
	_ = x[ReturnAttrNonNull-3]
	_ = x[ReturnAttrNoUndef-4]
	_ = x[ReturnAttrNullPointerIsValid-5]
	_ = x[ReturnAttrSignExt-6]
	_ = x[ReturnAttrZeroExt-7]
}

const _ReturnAttr_name = "inregnoaliasnomergenonnullnoundefnull_pointer_is_validsignextzeroext"

var _ReturnAttr_index = [...]uint8{0, 5, 12, 19, 26, 33, 54, 61, 68}

func (i ReturnAttr) String() string {
	if i >= ReturnAttr(len(_ReturnAttr_index)-1) {
		return "ReturnAttr(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ReturnAttr_name[_ReturnAttr_index[i]:_ReturnAttr_index[i+1]]
}
