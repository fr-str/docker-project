// Code generated by "stringer -type=Major"; DO NOT EDIT.

package wsc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InternalServerError-100]
	_ = x[OK-200]
	_ = x[Error-300]
	_ = x[NotFound-400]
	_ = x[Missing-500]
	_ = x[Timeout-600]
	_ = x[Forbbiden-700]
	_ = x[Exists-800]
	_ = x[Invalid-900]
}

const (
	_Major_name_0 = "InternalServerError"
	_Major_name_1 = "OK"
	_Major_name_2 = "Error"
	_Major_name_3 = "NotFound"
	_Major_name_4 = "Missing"
	_Major_name_5 = "Timeout"
	_Major_name_6 = "Forbbiden"
	_Major_name_7 = "Exists"
	_Major_name_8 = "Invalid"
)

func (i Major) String() string {
	switch {
	case i == 100:
		return _Major_name_0
	case i == 200:
		return _Major_name_1
	case i == 300:
		return _Major_name_2
	case i == 400:
		return _Major_name_3
	case i == 500:
		return _Major_name_4
	case i == 600:
		return _Major_name_5
	case i == 700:
		return _Major_name_6
	case i == 800:
		return _Major_name_7
	case i == 900:
		return _Major_name_8
	default:
		return "Major(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}