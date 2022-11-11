// Code generated by "stringer -type=Type wsc/wsCodes.go"; DO NOT EDIT.

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
	_ = x[ReqID-0]
	_ = x[Action-1]
	_ = x[ActionArgs-2]
	_ = x[ContainerName-3]
	_ = x[Container-4]
	_ = x[ContainerIsRunning-5]
	_ = x[Decode-6]
	_ = x[Exited-7]
	_ = x[Inspect-8]
}

const (
	_Type_name_0 = "ReqIDActionActionArgsContainerNameContainerContainerIsRunningDecodeExitedInspect"
	_Type_name_1 = "InternalServerError"
	_Type_name_2 = "OK"
	_Type_name_3 = "Error"
	_Type_name_4 = "NotFound"
	_Type_name_5 = "Missing"
	_Type_name_6 = "Timeout"
	_Type_name_7 = "Forbbiden"
	_Type_name_8 = "Exists"
	_Type_name_9 = "Invalid"
)

var (
	_Type_index_0 = [...]uint8{0, 5, 11, 21, 34, 43, 61, 67, 73, 80}
)

func (i Type) String() string {
	switch {
	case i <= 8:
		return _Type_name_0[_Type_index_0[i]:_Type_index_0[i+1]]
	case i == 100:
		return _Type_name_1
	case i == 200:
		return _Type_name_2
	case i == 300:
		return _Type_name_3
	case i == 400:
		return _Type_name_4
	case i == 500:
		return _Type_name_5
	case i == 600:
		return _Type_name_6
	case i == 700:
		return _Type_name_7
	case i == 800:
		return _Type_name_8
	case i == 900:
		return _Type_name_9
	default:
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}